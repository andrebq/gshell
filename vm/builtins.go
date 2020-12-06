package vm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/ast/match"
)

var (
	StdoutChannel = mustSym("stdout")
	StdinChannel  = mustSym("stdin")
	StderrChannel = mustSym("stderr")
)

func GShellLetVariable(c *CallStack) {
	if len(c.RawArgs) != 2 {
		c.FailWith = errors.New("Invalid use of 'let', should be: 'let <$variableName> <value>'")
		return
	}
	varname := c.RawArgs[0].(ast.Var).Name()
	parent := c.Context.parent
	if !parent.CanBind(varname) {
		c.FailWith = fmt.Errorf("Variable %v is already defined in the context", varname.Text())
		return
	}
	val, err := c.VM.Eval(c.Context, c.RawArgs[1])
	if err != nil {
		c.FailWith = err
		return
	}
	parent.Let(varname, val)
	c.ReturnValue = val
}

func GShellPrintln(c *CallStack) {
	parts := make([]string, len(c.RawArgs))
	for i, a := range c.RawArgs {
		v, err := c.VM.Eval(c.Context, a)
		if err != nil {
			c.FailWith = err
			return
		}
		switch v := v.(type) {
		case ast.Symbol:
			parts[i] = v.Text()
		case ast.Number:
			parts[i] = strconv.FormatFloat(v.Float64(), 'f', -1, 64)
		default:
			parts[i] = fmt.Sprintf("%v", v)
		}
	}
	c.VM.PushValues(StdoutChannel, c.Context, strings.Join(parts, " ")+"\n")
	c.ReturnValue = trueSym
}

func GShellIf(c *CallStack) {
	switch len(c.RawArgs) {
	case 2:
	case 4:
		break
	default:
		c.FailWith = errors.New("If command: if <expression> { <commands > } [else { <commands> } ]")
		return
	}
	var exp ast.Argument
	if !match.Apply(&c.RawArgs, match.Head(&exp)) {
		c.FailWith = errors.New("Missing conditional expression")
	}
	var thenBlock *ast.Script
	if !match.Apply(&c.RawArgs, match.Script(&thenBlock)) {
		c.FailWith = errors.New("Missing then block")
		return
	}

	var elseBlock *ast.Script
	match.Apply(&c.RawArgs, match.Guard(match.Symbol(elseSym), match.Script(&elseBlock)))

	cond, err := c.VM.Eval(c.Context, exp)
	if err != nil {
		c.FailWith = err
		return
	}

	var boolValue bool
	c.FailWith = c.VM.CastTo(c.Context, cond, &boolValue)
	if c.FailWith != nil {
		return
	}
	if boolValue {
		c.ReturnValue, c.FailWith = c.VM.Eval(c.Context, thenBlock)
	} else if elseBlock != nil {
		c.ReturnValue, c.FailWith = c.VM.Eval(c.Context, elseBlock)
	}
}
