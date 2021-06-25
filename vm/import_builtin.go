package vm

import (
	"errors"
	"fmt"

	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/ast/match"
)

func GShellImport(c *CallStack) {
	importUsage := `import [ <["module/path/or/Identifier" local-name] ...>]`
	var lst *ast.List
	type mapping struct {
		Path ast.Text
		Sym  ast.Symbol
	}
	mappings := []mapping{}
	if !match.Apply(&c.RawArgs, match.ListOf(&lst, func(a ast.Argument) bool {
		moduleDef, ok := a.(*ast.List)
		if !ok {
			return false
		}
		m := mapping{}
		vec := moduleDef.ToSlice(nil)
		ok = match.Apply(&vec, match.Guard(match.Text(&m.Path), match.AnySymbol(&m.Sym))) &&
			len(vec) == 0
		if !ok {
			return false
		}
		mappings = append(mappings, m)
		return ok
	})) {
		c.FailWith = errors.New(importUsage)
	}
	for _, m := range mappings {
		err := c.VM.loadModule(c.Context, m.Sym, m.Path)
		if err != nil {
			// should we keep this information private?!
			c.FailWith = fmt.Errorf("Error while loading %v as %v, cause: %v", m.Path, m.Sym, err)
			return
		}
	}
	c.ReturnValue = trueSym
}
