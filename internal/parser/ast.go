package parser

import (
	"bytes"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type (
	Ast struct {
		root *Script
	}

	stack []interface{}
	queue []interface{}

	Cmd struct {
		command Symbol
		args    []interface{}
	}

	Script struct {
		cmds []*Cmd
	}

	Text struct {
		string
	}

	Number struct {
		float64
	}

	Symbol struct {
		sym string
	}

	Formatter interface {
		Fmt(Printer)
	}

	astBuilder struct {
		*BaseGShellListener
		ast *Ast

		stack stack
	}
)

func (s Symbol) Fmt(p Printer) {
	p.WriteString(s.sym)
}

func (n Number) Fmt(p Printer) {
	p.WriteString(strconv.FormatFloat(n.float64, 'f', -1, 64))
}

func (t Text) Fmt(p Printer) {
	p.WriteString(t.string)
}

func (c *Cmd) Fmt(p Printer) {
	c.command.Fmt(p)
	for _, a := range c.args {
		if f, ok := a.(Formatter); ok {
			p.WriteArgSeparator()
			f.Fmt(p)
		}
	}
}

func (s *Script) Fmt(p Printer) {
	p.WriteIndent()
	p.WriteString("{")
	switch len(s.cmds) {
	case 0:
		p.WriteString("}")
		return
	case 1:
		p.WriteString(" ")
		s.cmds[0].Fmt(p)
		p.WriteInlineTerminator()
		p.WriteString(" }")
		return
	default:
		p.WriteLineBreak()
		p.Indent()

		for _, c := range s.cmds {
			p.WriteIndent()
			c.Fmt(p)
			p.WriteLineBreak()
		}
		p.Unindent()
		p.WriteIndent()
		p.WriteString("}")
		return
	}
}

func (a *Ast) Fmt(p Printer) {
	if a.root == nil {
		p.WriteString("{}")
		return
	}
	a.root.Fmt(p)
}

func (a *Ast) String() string {
	buf := &bytes.Buffer{}
	p := &printer{
		out: buf,
	}
	a.Fmt(p)
	if p.Err() != nil {
		return "!invalid-ast"
	}
	return buf.String()
}

func newAstBuilder() *astBuilder {
	return &astBuilder{
		BaseGShellListener: &BaseGShellListener{},
		ast:                &Ast{},
	}
}

func Parse(code string) (*Ast, error) {
	lexer := NewGShellLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGShellParser(stream)
	errorsFound := &errorsListener{
		ErrorListener: antlr.NewDefaultErrorListener(),
	}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorsFound)
	astBuilder := newAstBuilder()
	antlr.ParseTreeWalkerDefault.Walk(astBuilder, parser.Start())

	if errorsFound.err != nil {
		return nil, errorsFound.err
	}
	return astBuilder.ast, nil
}

func (ab *astBuilder) ExitStart(c *StartContext) {
	ab.ast = &Ast{root: ab.stack.pop().(*Script)}
}

func (ab *astBuilder) EnterScript(c *ScriptContext) {
	ab.stack.push(&Script{})
}

func (ab *astBuilder) ExitScript(c *ScriptContext) {
	steps := new(stack)
	for !isScript(ab.stack) {
		steps.push(ab.stack.pop().(*Cmd))
	}

	sc := ab.stack.pop().(*Script)
	sc.cmds = make([]*Cmd, 0, steps.depth())
	for !steps.empty() {
		sc.cmds = append(sc.cmds, steps.pop().(*Cmd))
	}
	ab.stack.push(sc)
}

func (ab *astBuilder) EnterSingleCommand(c *SingleCommandContext) {
	ab.stack.push(&Cmd{})
}

func (ab *astBuilder) ExitSingleCommand(c *SingleCommandContext) {
	argst := new(stack)
	for !isCmd(ab.stack) {
		argst.push(ab.stack.pop())
	}

	cmd := ab.stack.pop().(*Cmd)

	cmd.args = make([]interface{}, 0, argst.depth())
	for !argst.empty() {
		if cmd.command == (Symbol{}) {
			cmd.command = argst.pop().(Symbol)
			continue
		}
		cmd.args = append(cmd.args, argst.pop())
	}

	ab.stack.push(cmd)
}

func (ab *astBuilder) ExitCommandName(c *CommandNameContext) {
	ab.stack.push(Symbol{sym: c.GetText()})
}

func (ab *astBuilder) ExitNumericArgument(c *NumericArgumentContext) {
	number, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		ab.stack.push(Text{c.GetText()})
	} else {
		ab.stack.push(Number{number})
	}
}

func (ab *astBuilder) ExitNamedArgument(c *NamedArgumentContext) {
	ab.stack.push(Symbol{sym: c.GetText()})
}

func (s *stack) push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) pop() interface{} {
	v := (*s)[len(*s)-1]
	// ensure gc cannot reach the old value
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) peek() interface{} {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool { return len(*s) == 0 }
func (s *stack) depth() int  { return len(*s) }

func isCmd(st stack) bool {
	_, ok := st.peek().(*Cmd)
	return ok
}

func isScript(st stack) bool {
	_, ok := st.peek().(*Script)
	return ok
}
