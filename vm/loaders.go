package vm

import (
	"context"
	"fmt"

	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/internal/parser"
)

type (
	memoryLoader struct {
		items map[ast.Text]ast.Ast
	}

	MemoryLoader interface {
		ModuleParser
		AddCode(path ast.Text, code string) error
		AddAst(path ast.Text, tree ast.Ast)
	}
)

func NewMemoryLoader() MemoryLoader {
	return &memoryLoader{
		items: make(map[ast.Text]ast.Ast),
	}
}

func (m *memoryLoader) AddAst(path ast.Text, tree ast.Ast) {
	m.items[path] = tree
}

func (m *memoryLoader) AddCode(path ast.Text, code string) error {
	script, err := parser.Parse(code)
	if err != nil {
		return err
	}
	m.items[path] = (*script)
	return nil
}

func (m *memoryLoader) Parse(ctx context.Context, path ast.Text) (ast.Ast, error) {
	p, ok := m.items[path]
	if !ok {
		return ast.Ast{}, fmt.Errorf("module %v not found", path)
	}
	return p, nil
}
