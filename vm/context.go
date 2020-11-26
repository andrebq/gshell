package vm

import "github.com/andrebq/gshell/ast"

func NewContext(parent *Context) *Context {
	return &Context{
		parent: parent,
		refs:   make(map[ast.Symbol]Value),
	}
}

func (c *Context) Set(sym ast.Symbol, v Value) {
	c.refs[sym] = v
}

func (c *Context) Get(sym ast.Symbol) (Value, bool) {
	v, ok := c.refs[sym]
	if !ok && c.parent != nil {
		return c.parent.Get(sym)
	}
	return v, ok
}
