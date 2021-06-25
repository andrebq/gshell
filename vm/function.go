package vm

import "github.com/andrebq/gshell/ast"

type (
	function struct {
		module   ast.Text
		upvalues *Context
		name     ast.Symbol
		args     []ast.Var
		body     *ast.Script
	}
)
