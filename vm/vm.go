package vm

import (
	"errors"
	"fmt"
	"sync"

	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/internal/parser"
)

type (
	VM struct {
		builtins map[ast.Symbol]Process

		modules map[ast.Symbol]*Module

		chlock        sync.RWMutex
		rootCtx       *Context
		currentModule ast.Symbol
	}

	Module struct {
		definitions *Context
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
	printlnSym = mustSym("println")
	letSym     = mustSym("let")
	switchSym  = mustSym("switch")
	elseSym    = mustSym("else")
	caseSym    = mustSym("case")
	guardSym   = mustSym("guard")
	fromSym    = mustSym("from")
	toSym      = mustSym("to")
	loop       = mustSym("loop")
	funcSym    = mustSym("func")

	trueSym  = mustSym("true")
	falseSym = mustSym("false")

	mainModuleSym = mustSym("main")

	ErrChannelNotFound = errors.New("channel not found")
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
	}
}

func NewRootContext() *Context {
	rootCtx := NewContext(nil)

	const defaultChanSize = 1000
	rootCtx.Set(StdoutChannel, make(chan Event, defaultChanSize))
	rootCtx.Set(StdinChannel, make(chan Event, defaultChanSize))
	rootCtx.Set(StdoutChannel, make(chan Event, defaultChanSize))

	return rootCtx
}

func NewVM() *VM {
	rootCtx := NewRootContext()

	vm := &VM{
		builtins: make(map[ast.Symbol]Process),
		modules: map[ast.Symbol]*Module{
			mainModuleSym: EmptyModule(rootCtx),
		},
		rootCtx:       rootCtx,
		currentModule: mainModuleSym,
	}
	vm.builtins[printlnSym] = ProcessFunc(GShellPrintln)
	vm.builtins[letSym] = ProcessFunc(GShellLetVariable)
	vm.builtins[switchSym] = ProcessFunc(GShellSwitch)
	vm.builtins[trueSym] = MakeIdentityProcess(trueSym)
	vm.builtins[falseSym] = MakeIdentityProcess(falseSym)
	vm.builtins[guardSym] = ProcessFunc(GShellGuard)
	vm.builtins[loop] = ProcessFunc(GShellLoop)
	vm.builtins[funcSym] = ProcessFunc(GShellFunc)

	return vm
}

func (v *VM) Stdout() <-chan Event {
	return v.getDefaultChannel(StdoutChannel)
}

func (v *VM) Stderr() <-chan Event {
	return v.getDefaultChannel(StderrChannel)
}

func (v *VM) Stdin() chan<- Event {
	return v.getDefaultChannel(StdinChannel)
}

func (v *VM) getDefaultChannel(name ast.Symbol) chan Event {
	ev, ok := v.rootCtx.Get(name)
	if !ok {
		return nil
	}
	return ev.(chan Event)
}

func (v *VM) Run(code string) (interface{}, error) {
	ast, err := parser.Parse(code)
	if err != nil {
		return nil, err
	}
	ctx := NewContext(v.rootCtx)
	return v.evalScript(ctx, ast.Root())
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
	v.evalScript(ctx, funcDeclaration.body)
}

func (v *VM) callValue(call *CallStack, value Value) {
	switch value := value.(type) {
	case Process:
		value.Run(call)
	default:
		call.FailWith = fmt.Errorf("Value %q is not callable", value)
	}
}

func (v *VM) callBuiltin(call *CallStack, value Process) {
	// things like for/if/let must be implemented as builtin
	value.Run(call)
}

func (v *VM) PushValues(chname ast.Symbol, ctx *Context, values ...Value) bool {
	ch, ok := ctx.Get(chname)
	if !ok {
		return false
	}
	if ch, ok := ch.(chan Event); ok {
		for _, v := range values {
			// TODO: this is blocking, allow some for of safely closing the VM
			// even when this operation is going on
			var ev Event
			switch v := v.(type) {
			case Event:
			default:
				ev.Main = v
			}
			ch <- ev
		}
		return true
	}
	return false
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

func (v *VM) newFunction(ctx *Context, module, name ast.Symbol, args *ast.List, script *ast.Script) *function {
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
