package vm

import "github.com/andrebq/gshell/ast"

type (
	function struct {
		module   ast.Symbol
		upvalues *Context
		name     ast.Symbol
		args     *ast.List
		body     *ast.Script
	}
)
