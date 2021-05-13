package vm

import "github.com/andrebq/gshell/ast"

func NewContext(parent *Context) *Context {
	return &Context{
		parent:     parent,
		refs:       make(map[ast.Symbol]Value),
		isFunction: parent.IsFunction(),
	}
}

func NewFunctionContext(parent *Context) *Context {
	return &Context{
		parent:     parent,
		refs:       make(map[ast.Symbol]Value),
		isFunction: true,
	}
}

func (c *Context) IsFunction() bool {
	if c == nil {
		return false
	}
	return c.isFunction || c.parent.IsFunction()
}

func (c *Context) Set(sym ast.Symbol, v Value) {
	c.refs[sym] = v
}

// Let sym be bound to v if, and only if, sym
// is not already defined in this context.
//
// Symbols from parent contexts can be overwritten though
func (c *Context) Let(sym ast.Symbol, v Value) bool {
	if !c.CanBind(sym) {
		return false
	}
	c.refs[sym] = v
	return true
}

func (c *Context) Get(sym ast.Symbol) (Value, bool) {
	v, ok := c.refs[sym]
	if !ok && c.parent != nil {
		return c.parent.Get(sym)
	}
	return v, ok
}

func (c *Context) CanBind(sym ast.Symbol) bool {
	_, ok := c.refs[sym]
	return !ok
}
