package vm

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/internal/parser"
	"github.com/andrebq/gshell/mailbox"
)

type (
	VM struct {
		builtins map[ast.Symbol]Process

		modules map[ast.Text]*Module

		loader ModuleParser

		chlock        sync.RWMutex
		rootCtx       *Context
		currentModule ast.Text

		// shortcuts to the actors registered in pid
		stdout *Actor
		stderr *Actor
		stdin  *Actor

		pids map[ast.Symbol]*Actor
	}

	ModuleParser interface {
		Parse(ctx context.Context, path ast.Text) (ast.Ast, error)
	}

	Actor struct {
		box        mailbox.Box
		scope      ast.Symbol
		identifier ast.Symbol
	}

	Module struct {
		definitions *Context
		loading     bool
	}

	Value interface{}

	Event struct {
		Main Value
		Meta map[string][]string
	}

	CallStack struct {
		RawArgs []ast.Argument
		VM      *VM
		Context *Context

		// Use to indcate that the function call failed
		// and give it a reason
		FailWith error

		ReturnValue interface{}
	}

	Context struct {
		goCtx      context.Context
		parent     *Context
		refs       map[ast.Symbol]Value
		isFunction bool
	}

	Process interface {
		Run(*CallStack)
	}

	ProcessFunc func(*CallStack)
)

var (
	printlnSym = ast.MustNewSymbol("println")
	letSym     = ast.MustNewSymbol("let")
	switchSym  = ast.MustNewSymbol("switch")
	elseSym    = ast.MustNewSymbol("else")
	caseSym    = ast.MustNewSymbol("case")
	guardSym   = ast.MustNewSymbol("guard")
	fromSym    = ast.MustNewSymbol("from")
	toSym      = ast.MustNewSymbol("to")
	loop       = ast.MustNewSymbol("loop")
	funcSym    = ast.MustNewSymbol("func")

	trueSym   = ast.MustNewSymbol("true")
	falseSym  = ast.MustNewSymbol("false")
	importSym = ast.MustNewSymbol("import")

	localSym = ast.MustNewSymbol("local")
	refSym   = ast.MustNewSymbol("refs")
	pidSym   = ast.MustNewSymbol("pids")

	stdoutSym = ast.MustNewSymbol("stdout")
	stderrSym = ast.MustNewSymbol("stderr")
	stdinSym  = ast.MustNewSymbol("stdin")

	mainModuleText = ast.NewText("<main.gshell>")

	ErrChannelNotFound = errors.New("channel not found")

	reservedScopes = []ast.Symbol{
		localSym,
		refSym,
		pidSym,
	}
)

func (pf ProcessFunc) Run(c *CallStack) {
	pf(c)
}

func mustSym(s string) ast.Symbol {
	sym, err := ast.NewSymbol(s)
	if err != nil {
		panic(err)
	}
	return sym
}

func EmptyModule(ctx *Context) *Module {
	return &Module{
		definitions: NewContext(ctx),
		loading:     true,
	}
}

func newRootContext() *Context {
	rootCtx := NewContext(nil)

	const defaultChanSize = 1000
	rootCtx.Set(stdinSym, make(chan Event, defaultChanSize))
	rootCtx.Set(stdoutSym, make(chan Event, defaultChanSize))
	rootCtx.Set(stderrSym, make(chan Event, defaultChanSize))

	return rootCtx
}

func NewVM() *VM {
	rootCtx := newRootContext()

	vm := &VM{
		builtins: make(map[ast.Symbol]Process),
		modules: map[ast.Text]*Module{
			mainModuleText: EmptyModule(rootCtx),
		},
		rootCtx:       rootCtx,
		currentModule: mainModuleText,
		pids:          make(map[ast.Symbol]*Actor),
	}

	vm.builtins[trueSym] = MakeIdentityProcess(trueSym)
	vm.builtins[falseSym] = MakeIdentityProcess(falseSym)
	vm.builtins[letSym] = ProcessFunc(GShellLetVariable)
	vm.builtins[funcSym] = ProcessFunc(GShellFunc)
	vm.builtins[importSym] = ProcessFunc(GShellImport)

	vm.builtins[printlnSym] = ProcessFunc(GShellPrintln)
	vm.builtins[switchSym] = ProcessFunc(GShellSwitch)
	vm.builtins[guardSym] = ProcessFunc(GShellGuard)
	vm.builtins[loop] = ProcessFunc(GShellLoop)

	vm.stdout = vm.newActor(localSym, stdoutSym)
	vm.stderr = vm.newActor(localSym, stderrSym)
	vm.stdin = vm.newActor(localSym, stdinSym)

	return vm
}

func (v *VM) SetModuleParser(p ModuleParser) {
	v.loader = p
}

func (v *VM) newActor(scope ast.Symbol, id ast.Symbol) *Actor {
	upid := ast.ScopedSymbol(scope, id)
	actor := v.pids[upid]
	if actor == nil {
		actor = &Actor{
			box:        mailbox.NonBlocking(100),
			scope:      scope,
			identifier: id,
		}
		v.pids[upid] = actor
	}
	return actor
}

func (v *VM) Stdout() mailbox.Reader {
	return v.stdout.box.Reader()
}

func (v *VM) Stderr() mailbox.Reader {
	return v.stderr.box.Reader()
}

func (v *VM) Run(code string) (interface{}, error) {
	ast, err := parser.Parse(code)
	if err != nil {
		return nil, err
	}
	return v.runModule(v.currentModule, ast)
}

func (v *VM) runModule(name ast.Text, ast *ast.Ast) (interface{}, error) {
	return v.switchModule(name, func(module *Module) (interface{}, error) {
		return v.evalScript(module.definitions, ast.Root())
	})
}

func (v *VM) switchModule(name ast.Text, fn func(*Module) (interface{}, error)) (interface{}, error) {
	oldModule := v.currentModule
	defer func() {
		v.currentModule = oldModule
	}()
	v.currentModule = name
	module, ok := v.modules[name]
	if !ok {
		module = EmptyModule(v.rootCtx)
		v.modules[name] = module
	}
	return fn(module)
}

func (v *VM) evalList(ctx *Context, lst *ast.List) (*ast.List, error) {
	output := ast.NilList()
	var err error
	lst.ForEach(func(a ast.Argument) bool {
		var value Value
		value, err = v.Eval(ctx, a)
		if err != nil {
			return false
		}
		var arg ast.Argument
		arg, err = ast.EncodeValue(value)
		if err != nil {
			return false
		}
		// this is very slow but I'm too lazy
		// to use a slice
		output = output.Append(arg)
		return true
	})
	return output, err
}

func (v *VM) evalScript(ctx *Context, sc *ast.Script) (Value, error) {
	var lastReturn Value
	for _, c := range sc.Commands() {
		call := v.runCommand(ctx, c)
		lastReturn = call.ReturnValue
		if call.FailWith != nil {
			// println
			return nil, call.FailWith
		}
	}
	return lastReturn, nil
}

func (v *VM) runCommand(ctx *Context, cmd *ast.Cmd) *CallStack {
	call := &CallStack{
		VM:      v,
		RawArgs: cmd.Arguments(),
		Context: ctx,
	}
	bt, found := v.builtins[cmd.Command()]
	if found {
		v.callBuiltin(call, bt)
		return call
	}
	value, found := ctx.Get(cmd.Command())
	if found {
		v.callValue(call, value)
		return call
	}

	// TODO: CONTINUE HERE
	// add a function to check if the cmd symbol is scoped
	// and extracts the scope, that will allow us to lookup the module
	//
	// rename local/refs/pids sccopes to _local/_refs/_pids scopes, that way
	// users are free to use that scope for local variables

	moduleFunc, found := v.modules[v.currentModule].definitions.Get(cmd.Command())
	if found {
		v.callFunction(call, moduleFunc.(*function))
		return call
	}

	call.FailWith = fmt.Errorf("Command %v not found", cmd.Command().Text())
	return call
}

func (v *VM) callFunction(call *CallStack, funcDeclaration *function) {
	ctx := NewFunctionContext(funcDeclaration.upvalues)
	if len(funcDeclaration.args) != len(call.RawArgs) {
		call.FailWith = fmt.Errorf("Function %v requires %v args got %v", funcDeclaration.name, len(funcDeclaration.args), len(call.RawArgs))
		return
	}
	for i := range call.RawArgs {
		var err error
		var argValue Value
		argValue, err = v.Eval(call.Context, call.RawArgs[i])
		if err != nil {
			call.FailWith = err
			return
		}
		ctx.Set(funcDeclaration.args[i].Name(), argValue)
	}
	call.ReturnValue, call.FailWith = v.evalScript(ctx, funcDeclaration.body)
}

func (v *VM) callValue(call *CallStack, value Value) {
	switch value := value.(type) {
	case Process:
		value.Run(call)
	case *function:
		v.callFunction(call, value)
	default:
		call.FailWith = fmt.Errorf("Value %q is not callable", value)
	}
}

func (v *VM) callBuiltin(call *CallStack, value Process) {
	// things like for/if/let must be implemented as builtin
	value.Run(call)
}

func (v *VM) enqueueValue(pid ast.Symbol, ctx *Context, values ...Value) bool {
	actor := v.pids[pid]
	if actor == nil {
		return false
	}
	for _, v := range values {
		err := mailbox.BlockingPush(ctx.goCtx, actor.box, v)
		if err != nil {
			// TODO: think about how to handle errors here
			return false
		}
	}
	return true
}

func (v *VM) Eval(ctx *Context, a ast.Argument) (interface{}, error) {
	switch a := a.(type) {
	case ast.Symbol:
		// a symbol always evaluates to itself
		return a, nil
	case ast.Number:
		return a.Float64(), nil
	case ast.Text:
		return a.Text(), nil
	case ast.Var:
		v, ok := ctx.Get(a.Name())
		if !ok {
			return nil, fmt.Errorf("Variable %v is not defined", a)
		}
		return v, nil
	case *ast.Script:
		return v.evalScript(ctx, a)
	case *ast.List:
		return v.evalList(ctx, a)
	}
	return nil, fmt.Errorf("cannot decode %T into a meangingful value", a)
}

func (v *VM) CastTo(ctx *Context, input interface{}, out interface{}) error {
	switch out := out.(type) {
	case *bool:
		return v.castToBool(ctx, input, out)
	case *float64:
		return v.castToFloat(ctx, input, out)
	}
	return fmt.Errorf("cast not allowed %T", out)
}

func (v *VM) EvalAndCast(ctx *Context, arg ast.Argument, out interface{}) error {
	value, err := v.Eval(ctx, arg)
	if err != nil {
		return err
	}
	return v.CastTo(ctx, value, out)
}

func (v *VM) loadModule(ctx *Context, dest ast.Symbol, path ast.Text) error {
	if v.loader == nil {
		return errors.New("there are no loaders configured for this VM")
	}

	if !ctx.CanBind(dest) {
		return fmt.Errorf("symbol %v already defined", dest)
	}

	if _, found := v.modules[path]; found {
		return nil
	}

	ast, err := v.loader.Parse(context.TODO(), path)
	if err != nil {
		return fmt.Errorf("unable to parse module %v, cause: %v", path, err)
	}
	_, err = v.runModule(path, &ast)
	return err
}

func (v *VM) castToBool(ctx *Context, input interface{}, out *bool) error {
	*out = false
	switch input := input.(type) {
	case ast.Symbol:
		*out = input == trueSym
	default:
		return fmt.Errorf("Cannot cast object of type %T to bool", input)
	}
	return nil
}

func (v *VM) castToFloat(ctx *Context, input interface{}, out *float64) error {
	*out = 0
	switch input := input.(type) {
	case ast.Number:
		*out = input.Float64()
	case float64:
		*out = input
	case int64:
		*out = float64(input)
	case float32:
		*out = float64(input)
	case int32:
		*out = float64(input)
	case int:
		*out = float64(input)
	default:
		return fmt.Errorf("Cannot cast object of type %T to float64", input)
	}
	return nil
}

func (v *VM) newFunction(ctx *Context, module ast.Text, name ast.Symbol, args *ast.List, script *ast.Script) *function {
	var items []ast.Var
	args.ForEach(func(a ast.Argument) bool {
		items = append(items, a.(ast.Var))
		return true
	})
	return &function{
		module:   module,
		upvalues: ctx,
		args:     items,
		body:     script,
	}
}
