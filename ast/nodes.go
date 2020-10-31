package ast

import (
	"bytes"
	"fmt"
	"io"

	"github.com/andrebq/gshell/lexer"
)

type (
	WriteString interface {
		io.Writer
		WriteString(str string) (int, error)
	}

	Node interface {
		// anchor just to avoid external types to be nodes
		// TODO: maybe ast could be an internal object
		anchor()

		writeTo(out WriteString, ident string)
	}

	baseNode struct{}

	Program struct {
		baseNode
		Instructions []*Pipeline
	}

	Pipeline struct {
		baseNode
		Items []*Command
	}

	Command struct {
		baseNode
		Identifier *Identifier
		Arguments  []Node
	}

	Identifier struct {
		baseNode
		Name lexer.Lexem
	}

	Number struct {
		baseNode
		Value lexer.Lexem
	}

	QuotedText struct {
		baseNode
		Text lexer.Lexem
	}
)

func (bn baseNode) anchor() {}

func (p *Pipeline) String() string {
	buf := &bytes.Buffer{}
	p.writeTo(buf, "")
	return buf.String()
}

func (p *Pipeline) writeTo(out WriteString, ident string) {
	out.WriteString(ident)
	nextIdent := ident + ident
	for i, cmd := range p.Items {
		if i != 0 {
			out.WriteString(" |")
		}
		out.WriteString("\n")
		cmd.writeTo(out, nextIdent)
	}
	out.WriteString("\n")
	out.WriteString(ident)
}

func (p *Command) writeTo(out WriteString, ident string) {
	out.WriteString(ident)
	out.WriteString(p.Identifier.Name.Value)
	for _, arg := range p.Arguments {
		out.WriteString(" ")
		arg.writeTo(out, "")
	}
}

func (i *Identifier) writeTo(out WriteString, ident string) {
	out.WriteString(ident)
	out.WriteString(fmt.Sprintf("%v", i.Name.Value))
}

func (n *Number) writeTo(out WriteString, ident string) {
	out.WriteString(ident)
	out.WriteString(fmt.Sprintf("%v", n.Value.Value))
}

func (q *QuotedText) writeTo(out WriteString, ident string) {
	out.WriteString(ident)
	out.WriteString(fmt.Sprintf("%q", q.Text.Value))
}
