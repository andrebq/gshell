// Package match provides helper functions to match arguments against a set of
// constratins
package match

import "github.com/andrebq/gshell/ast"

type (
	Condition func(*[]ast.Argument) bool
)

// Apply the Condition to the given list of arguments
func Apply(args *[]ast.Argument, c Condition) bool {
	return c(args)
}

// Guard takes a list of Conditions and applies them sequentially
//
// If all of them are true the array is updated and the function
// returns true.
//
// The conditions are evaluated in sequence and the list of arguments
// is mutated in between the calls
func Guard(conditions ...Condition) Condition {
	return func(args *[]ast.Argument) bool {
		aux := *args
		for _, c := range conditions {
			if !c(&aux) {
				return false
			}
		}
		*args = aux
		return true
	}
}

// List returns the first ast.Argument as a List
func List(out **ast.List) Condition {
	return func(args *[]ast.Argument) bool {
		if len(*args) == 0 {
			return false
		}
		list, ok := (*args)[0].(*ast.List)
		if !ok {
			return false
		}
		*args = (*args)[1:]
		*out = list
		return true
	}
}

// ListOf returns the first ast.Argument as a List
// if, and only if, filter returns true for all elements
func ListOf(out **ast.List, condition func(ast.Argument) bool) Condition {
	return func(args *[]ast.Argument) bool {
		if len(*args) == 0 {
			return false
		}
		list, ok := (*args)[0].(*ast.List)
		if !ok {
			return false
		}
		valid := true
		list.ForEach(func(a ast.Argument) bool {
			valid = valid && condition(a)
			return valid
		})
		if !valid {
			return false
		}
		*args = (*args)[1:]
		*out = list
		return true
	}
}

// Head value from args and set its content to out,
// returns true if len(args) >= 0
//
// Args will be updated to contain the tail of the list
func Head(out *ast.Argument) Condition {
	return func(args *[]ast.Argument) bool {
		if len(*args) == 0 {
			return false
		}
		*out = (*args)[0]
		*args = (*args)[1:]
		return true
	}
}

// Symbol returns true if the first item in the list of arguments
// matches the expected symbol
func Symbol(sym ast.Symbol) Condition {
	return func(args *[]ast.Argument) bool {
		if len(*args) == 0 {
			return false
		}
		if (*args)[0] != sym {
			return false
		}
		*args = (*args)[1:]
		return true
	}
}

// Script takes the first item from the list if it is an script block
func Script(out **ast.Script) Condition {
	return func(args *[]ast.Argument) bool {
		if len(*args) == 0 {
			return false
		}
		aux, ok := (*args)[0].(*ast.Script)
		if !ok {
			return false
		}
		*out = aux
		*args = (*args)[1:]
		return true
	}
}

// AnySymbol matches against anything that is a Symbol
func AnySymbol(out *ast.Symbol) Condition {
	return func(args *[]ast.Argument) bool {
		if len(*args) == 0 {
			return false
		}
		aux, ok := (*args)[0].(ast.Symbol)
		if !ok {
			return false
		}
		*out = aux
		*args = (*args)[1:]
		return true
	}
}

func Text(out *ast.Text) Condition {
	return func(args *[]ast.Argument) bool {
		if len(*args) == 0 {
			return false
		}
		var ok bool
		*out, ok = (*args)[0].(ast.Text)
		if !ok {
			return false
		}
		*args = (*args)[1:]
		return true
	}
}
