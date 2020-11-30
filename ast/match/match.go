// Package match provides helper functions to match arguments against a set of
// constratins
package match

import "github.com/andrebq/gshell/ast"

// Head value from args and set its content to out,
// returns true if len(args) >= 0
//
// Args will be updated to contain the tail of the list
func Head(out *ast.Argument, args *[]ast.Argument) bool {
	if len(*args) == 0 {
		return false
	}
	*out = (*args)[0]
	*args = (*args)[1:]
	return true
}
