package shell

import (
	"errors"
	"fmt"

	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/lexer"
	"github.com/andrebq/gshell/parser"
)

type (
	// VM contains the code for all modules loaded and the current execution environment
	VM struct {
		programs    map[string]*Program
		environment map[string]string
		builtin     map[string]Builtin
	}

	// CallContext contains the information that a VM
	// can provide for a Program/Builtin
	CallContext struct {
		// Command contains the command identifier and all arguments
		//
		// In case of a Builtin program arguments might contain other programs,
		// in case of normal programs or other scripts, arguments are made up
		// of atomic units as those need to serialized.
		Command *ast.Command

		// VM contains the link to the underlying VM
		VM *VM

		// TODO(andre): add here the IO ports
	}

	// Builtin implements a internal VM command
	Builtin interface {
		Run(*CallContext) error
	}

	// BuiltinFunc implements Builtin for simple functions
	BuiltinFunc func(*CallContext) error

	Program struct {
		pipelines []*pipeline
	}

	pipeline struct {
		commands []*command
	}

	command struct {
		name string
		args []argument
	}

	// TODO: change this so we can handle things other than hard-coded strings
	argument struct {
		text string
	}
)

func NewVM() *VM {
	return &VM{
		programs:    make(map[string]*Program),
		environment: make(map[string]string),
	}
}

func (v *VM) ParseInteractive(name, code string) error {
	l, err := lexer.Lex([]byte(code))
	if err != nil {
		return err
	}
	pipeline, err := parser.ParseInteractive(l)
	if err != nil {
		return err
	}
	prog := NewProgram(pipeline)
	if err != nil {
		return err
	}
	v.programs[name] = prog
	return nil
}

func (v *VM) Run(name string) error {
	p, ok := v.programs[name]
	if !ok {
		return errors.New("program not found")
	}
	return p.runOn(v)
}

func NewProgram(items ...*ast.Pipeline) *Program {
	p := &Program{}
	for _, i := range items {
		p.pipelines = append(p.pipelines, newPipeline(i))
	}
	return p
}

func newPipeline(i *ast.Pipeline) *pipeline {
	p := &pipeline{}
	for _, cmd := range i.Items {
		p.commands = append(p.commands, newCommand(cmd))
	}
	return p
}

func newCommand(cmd *ast.Command) *command {
	c := &command{
		name: cmd.Identifier.Name.Value,
	}
	for _, arg := range cmd.Arguments {
		c.args = append(c.args, newArgument(arg))
	}
	return c
}

func newArgument(n ast.Node) argument {
	switch n := n.(type) {
	case *ast.QuotedText:
		return argument{n.Text.Value}
	case *ast.Number:
		return argument{n.Value.Value}
	case *ast.Identifier:
		return argument{n.Name.Value}
	}
	panic(fmt.Sprintf("cannot handle argument [%v]", n))
}

func (bf BuiltinFunc) Run(c *CallContext) error {
	return bf(c)
}
