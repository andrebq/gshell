package ast

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
)

type (
	Ast struct {
		root *Script
	}

	Cmd struct {
		command Symbol
		args    []Argument
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

	Var struct {
		sym Symbol
	}

	List struct {
		val  Argument
		tail *List
	}

	Formatter interface {
		Fmt(Printer)
	}

	Argument interface {
		anchor()
		Fmt(Printer)
	}
)

var (
	// this should be kept in sync with the parser rules
	// otherwise the package won't be able to create the Symbol node
	// TODO: think of a way to avoid this duplicated re
	symbolRe = regexp.MustCompile(`[\p{Ll}|\p{Lu}|!|?|\.|\\|-|+|*|&|^|%|#|@|~]+[\p{Ll}|\p{Lu}|!|?|\.|\\|-|+|*|&|^|%|$|#|@|~|0-9]*`)

	errNotAValidSymbol = errors.New("not a valid symbol")
)

func (s Symbol) Fmt(p Printer) {
	p.WriteString(s.sym)
}
func (s Symbol) Text() string   { return s.sym }
func (s Symbol) anchor()        {}
func (s Symbol) String() string { return s.sym }

func (n Number) Fmt(p Printer) {
	p.WriteString(strconv.FormatFloat(n.float64, 'f', -1, 64))
}
func (n Number) Float64() float64 { return n.float64 }
func (n Number) anchor()          {}
func (n Number) String() string   { return strconv.FormatFloat(n.float64, 'f', -1, 64) }

func (t Text) Fmt(p Printer) {
	p.WriteString(t.string)
}
func (t Text) Text() string   { return t.string }
func (t Text) String() string { return t.string }
func (t Text) anchor()        {}

func (v Var) anchor() {}
func (v Var) Fmt(p Printer) {
	p.WriteString("$")
	v.sym.Fmt(p)
}
func (v Var) Name() Symbol   { return v.sym }
func (v Var) String() string { return "$" + v.sym.Text() }

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

func (s *Script) anchor() {}

func (a *Ast) Fmt(p Printer) {
	if a.root == nil {
		p.WriteString("{}")
		return
	}
	a.root.Fmt(p)
}

func (a *Ast) Root() *Script { return a.root }

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

func (a *Ast) SetRoot(sc *Script) *Script {
	a.root = sc
	return a.root
}

func New() *Ast {
	return &Ast{
		root: NewScript(),
	}
}

func NewScript() *Script {
	return &Script{}
}

func NewCmd() *Cmd {
	return &Cmd{}
}

func NewNumber(v float64) Number {
	return Number{float64: v}
}

func NewText(t string) Text {
	return Text{string: t}
}

func NewSymbol(s string) (Symbol, error) {
	if !symbolRe.MatchString(s) {
		return Symbol{}, errNotAValidSymbol
	}
	return Symbol{sym: s}, nil
}

func NewVar(sym Symbol) Var { return Var{sym} }
func NewVarString(s string) (Var, error) {
	sym, err := NewSymbol(s)
	return NewVar(sym), err
}

func (s *Script) AddCommand(cmd *Cmd) *Script {
	s.cmds = append(s.cmds, cmd)
	return s
}

func (s *Script) Commands() []*Cmd {
	return append([]*Cmd(nil), s.cmds...)
}

func (c *Cmd) SetCommand(s Symbol) *Cmd {
	c.command = s
	return c
}

func (c *Cmd) Command() Symbol {
	return c.command
}

func (c *Cmd) AddArgument(a Argument) *Cmd {
	c.args = append(c.args, a)
	return c
}

func (c *Cmd) Arguments() []Argument {
	return append([]Argument(nil), c.args...)
}
