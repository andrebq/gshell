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

		chlock  sync.RWMutex
		rootCtx *Context
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
		parent *Context
		refs   map[ast.Symbol]Value
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

	trueSym  = mustSym("true")
	falseSym = mustSym("false")

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

func NewVM() *VM {
	vm := &VM{
		builtins: make(map[ast.Symbol]Process),
		rootCtx:  NewContext(nil),
	}
	vm.builtins[printlnSym] = ProcessFunc(GShellPrintln)
	vm.builtins[letSym] = ProcessFunc(GShellLetVariable)
	vm.builtins[switchSym] = ProcessFunc(GShellSwitch)
	vm.builtins[trueSym] = MakeIdentityProcess(trueSym)
	vm.builtins[falseSym] = MakeIdentityProcess(falseSym)
	vm.builtins[guardSym] = ProcessFunc(GShellGuard)

	const defaultChanSize = 1000
	vm.rootCtx.Set(StdoutChannel, make(chan Event, defaultChanSize))
	vm.rootCtx.Set(StdinChannel, make(chan Event, defaultChanSize))
	vm.rootCtx.Set(StdoutChannel, make(chan Event, defaultChanSize))
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
	return v.runScript(ctx, ast.Root())
}

func (v *VM) runScript(ctx *Context, sc *ast.Script) (interface{}, error) {
	var lastReturn interface{}
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
	if !found {
		call.FailWith = fmt.Errorf("Command %v not found", cmd.Command().Text())
		return call
	}
	v.callValue(call, value)
	return call
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
		return v.runScript(ctx, a)
	}
	return nil, fmt.Errorf("cannot decode %T into a meangingful value", a)
}

func (v *VM) CastTo(ctx *Context, input interface{}, out interface{}) error {
	switch out := out.(type) {
	case *bool:
		return v.castToBool(ctx, input, out)
	}
	return errors.New("cast not allowed")
}

func (v *VM) castToBool(ctx *Context, input interface{}, out *bool) error {
	*out = false
	switch input := input.(type) {
	case ast.Symbol:
		*out = input == trueSym
	default:
		return fmt.Errorf("Cannot cast object of type %t to bool", input)
	}
	return nil
}
