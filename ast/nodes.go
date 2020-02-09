package ast

import (
	"github.com/andrebq/gshell/lexer"
)

type (
	Node interface {
		// anchor just to avoid external types to be nodes
		// TODO: maybe ast could be an internal object
		anchor()
	}

	baseNode struct{}

	Pipeline struct {
		baseNode
		Items []*Command
	}

	Command struct {
		baseNode
		Identifier lexer.Lexem
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
