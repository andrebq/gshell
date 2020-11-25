package vm

import (
	"errors"
	"fmt"
	"sync"

	"github.com/andrebq/gshell/ast"
)

type (
	VM struct {
		builtins map[ast.Symbol]Process
		channels map[ast.Symbol]chan Value

		chlock sync.RWMutex
	}

	Value interface{}

	CallStack struct {
		RawArgs []ast.Argument
		VM      *VM
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
		channels: make(map[ast.Symbol]chan Value),
	}
	vm.builtins[printlnSym] = ProcessFunc(GShellPrintln)
	const defaultChanSize = 1000
	vm.channels[StdoutChannel] = make(chan Value, defaultChanSize)
	vm.channels[StdinChannel] = make(chan Value, defaultChanSize)
	vm.channels[StderrChannel] = make(chan Value, defaultChanSize)
	return vm
}

func (v *VM) GetChannel(chname ast.Symbol) (<-chan Value, error) {
	v.chlock.RLock()
	ch, ok := v.channels[chname]
	v.chlock.RUnlock()
	if !ok {
		return nil, ErrChannelNotFound
	}
	return ch, nil
}

func (v *VM) Run(code string) {
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
		ch <- v
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
