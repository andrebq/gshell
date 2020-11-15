package astutil

import (
	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/lexer"
)

func Pipeline(items ...*ast.Command) *ast.Pipeline {
	return &ast.Pipeline{
		Items: items,
	}
}

func Command(name *ast.Identifier, args ...ast.Node) *ast.Command {
	return &ast.Command{
		Identifier: name,
		Arguments:  args,
	}
}

func Identifier(name string) *ast.Identifier {
	return &ast.Identifier{
		Name: lexer.Lexem{Value: name, Type: lexer.Identifier},
	}
}

func QuotedText(text string) *ast.QuotedText {
	return &ast.QuotedText{
		Text: lexer.Lexem{Value: text, Type: lexer.QuotedText},
	}
}

func Number(value string) *ast.Number {
	return &ast.Number{
		Value: lexer.Lexem{Value: value, Type: lexer.Number},
	}
}
