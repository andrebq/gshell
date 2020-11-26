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
		channels map[ast.Symbol]chan Event

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
		Context Context
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
		channels: make(map[ast.Symbol]chan Event),
		rootCtx:  NewContext(nil),
	}
	vm.builtins[printlnSym] = ProcessFunc(GShellPrintln)
	const defaultChanSize = 1000
	vm.rootCtx.Set(StdoutChannel, make(chan Event, defaultChanSize))
	vm.rootCtx.Set(StdinChannel, make(chan Event, defaultChanSize))
	vm.rootCtx.Set(StdoutChannel, make(chan Event, defaultChanSize))
	return vm
}

func (v *VM) GetChannel(chname ast.Symbol) (<-chan Event, error) {
	v.chlock.RLock()
	ch, ok := v.channels[chname]
	v.chlock.RUnlock()
	if !ok {
		return nil, ErrChannelNotFound
	}
	return ch, nil
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
	return errors.New("not implemented")
}

func (v *VM) PushValues(chname ast.Symbol, values ...Value) bool {
	v.chlock.RLock()
	ch, ok := v.channels[chname]
	v.chlock.RUnlock()
	if !ok {
		return false
	}
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

func (v *VM) Eval(a ast.Argument) (interface{}, error) {
	switch a := a.(type) {
	case ast.Symbol:
		return a.Text(), nil
	case ast.Number:
		return a.Float64(), nil
	case ast.Text:
		return a.Text(), nil
	}
	return nil, fmt.Errorf("cannot decode %v into a meangingful value", a)
}
