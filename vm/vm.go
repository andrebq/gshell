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

func (v *VM) Run(code string) error {
	ast, err := parser.Parse(code)
	if err != nil {
		return err
	}
	ctx := NewContext(v.rootCtx)
	return v.runScript(ctx, ast.Root())
}

func (v *VM) runScript(ctx *Context, sc *ast.Script) error {
	for _, c := range sc.Commands() {
		err := v.runCommand(NewContext(ctx), c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VM) runCommand(ctx *Context, cmd *ast.Cmd) error {
	call := &CallStack{
		VM:      v,
		RawArgs: cmd.Arguments(),
		Context: ctx,
	}
	bt, found := v.builtins[cmd.Command()]
	if found {
		return v.callBuiltin(call, bt)
	}
	value, found := ctx.Get(cmd.Command())
	if !found {
		return fmt.Errorf("Command %v not found", cmd.Command().Text())
	}
	return v.callValue(call, value)
}

func (v *VM) callValue(call *CallStack, value Value) error {
	switch value := value.(type) {
	case Process:
		value.Run(call)
	default:
		return fmt.Errorf("Value %q is not callable", value)
	}
	return call.FailWith
}

func (v *VM) callBuiltin(call *CallStack, value Process) error {
	// things like for/if/let must be implemented as builtin
	value.Run(call)
	return call.FailWith
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
		return a.Text(), nil
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
	}
	return nil, fmt.Errorf("cannot decode %v into a meangingful value", a)
}
