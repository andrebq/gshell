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

func GShellGuard(c *CallStack) {
	var cond *ast.Script
	var body *ast.Script
	if !match.Apply(&c.RawArgs, match.Guard(match.Script(&cond), match.Script(&body))) {
		c.FailWith = errors.New("guard {... conditions} {... actions}")
		return
	}

	condCtx := NewContext(c.Context)
	condEval, err := c.VM.runScript(condCtx, cond)
	if err != nil {
		c.FailWith = err
	}
	var allowed bool
	if err = c.VM.CastTo(condCtx, condEval, &allowed); err != nil {
		c.FailWith = err
	}
	if !allowed {
		c.ReturnValue = false
		return
	}

	bodyEval, err := c.VM.runScript(c.Context, body)
	if err != nil {
		c.FailWith = err
		return
	}
	c.ReturnValue = bodyEval
	return
}

func GShellSwitch(c *CallStack) {
	switchUseCase := "switch { case <<guard-script> <action-script> [<guard-script> <action-script>...] [else <action-script>] }"
	var clauses *ast.Script
	if !match.Apply(&c.RawArgs, match.Script(&clauses)) {
		c.FailWith = errors.New(switchUseCase)
		return
	}

	type switchCase struct {
		guard *ast.Script
		body  *ast.Script
	}

	var cases []switchCase

	for _, p := range clauses.Commands() {
		switch p.Command() {
		case caseSym:
			var guard *ast.Script
			var block *ast.Script
			caseArgs := p.Arguments()
			if !match.Apply(&caseArgs, match.Guard(match.Script(&guard), match.Script(&block))) {
				c.FailWith = errors.New(switchUseCase)
				return
			}
			cases = append(cases, switchCase{guard: guard, body: block})
		case elseSym:
			var block *ast.Script
			elseArgs := p.Arguments()
			if !match.Apply(&elseArgs, match.Script(&block)) {
				c.FailWith = errors.New(switchUseCase)
				return
			}
			cases = append(cases, switchCase{body: block})
		default:
			c.FailWith = errors.New(switchUseCase)
		}
	}

	for _, sc := range cases {
		ctx := NewContext(c.Context)
		var boolVal bool
		if sc.guard != nil {
			val, err := c.VM.Eval(ctx, sc.guard)
			if err != nil {
				c.FailWith = err
				return
			}
			if err := c.VM.CastTo(ctx, val, &boolVal); err != nil {
				c.FailWith = err
			}
		} else {
			boolVal = true
		}

		if boolVal {
			ctx = NewContext(c.Context)
			c.ReturnValue, c.FailWith = c.VM.Eval(ctx, sc.body)
			return
		}
	}
}

func GShellLoop(c *CallStack) {
	loopUsage := fmt.Sprintf("loop <variable-name> from <start-point> to <end-point> <{body}>")
	var varName ast.Symbol
	var fromArg, toArg ast.Argument
	var body *ast.Script
	guard := match.Guard(match.AnySymbol(&varName), match.Symbol(fromSym), match.Head(&fromArg),
		match.Symbol(toSym), match.Head(&toArg), match.Script(&body))
	if !match.Apply(&c.RawArgs, guard) {
		c.FailWith = errors.New(loopUsage)
		return
	}
	var fromIdx, toIdx float64
	c.FailWith = c.VM.EvalAndCast(c.Context, fromArg, &fromIdx)
	if c.FailWith != nil {
		return
	}
	c.FailWith = c.VM.EvalAndCast(c.Context, toArg, &toIdx)
	if c.FailWith != nil {
		return
	}
	rev := fromIdx > toIdx

	lst := ast.NilList()
	for (rev && fromIdx >= toIdx) || (!rev && fromIdx <= toIdx) {
		ctx := NewContext(c.Context)
		ctx.Set(varName, ast.NewNumber(fromIdx))
		val, err := c.VM.Eval(ctx, body)
		if err != nil {
			c.FailWith = err
			return
		}
		lst = lst.Append(val.(ast.Argument))

		if rev {
			fromIdx--
		} else {
			fromIdx++
		}
	}
	c.ReturnValue = lst
}

func GShellFunc(c *CallStack) {
	// for now, let's just now allow functions to define functions
	// this is cutting on the capability of creating closures
	// which is kind of important on a purely immutable language...
	//
	// but bash doesn't have it and it kinda works ok there :-)
	// so no fancy schmancy high-order functions for you!
	if c.Context.isFunction {
		c.FailWith = errors.New("Functions cannot define other functions... sorry :(")
		return
	}
	funcUsage := fmt.Sprintf("func <function-name> [<variable list>]  <{function body}>")
	var funcName ast.Symbol
	var argList *ast.List
	var body *ast.Script
	guard := match.Guard(match.AnySymbol(&funcName),
		match.List(&argList),
		match.Script(&body))
	if !match.Apply(&c.RawArgs, guard) {
		c.FailWith = errors.New(funcUsage)
		return
	}

	cm := c.VM.currentModule
	module := c.VM.modules[c.VM.currentModule]
	if module == nil {
		c.FailWith = errors.New("VM does not have a current module therefore a function cannot be defined")
		return
	}

	if _, defined := module.definitions[funcName]; defined {
		c.FailWith = fmt.Errorf("Module %v already contains a function called %v.", cm, funcName)
		return
	}

	fn := c.VM.newFunction(c.Context, cm, funcName, argList, body)
	module.definitions[funcName] = fn
	c.ReturnValue = fn
	return
}

func MakeIdentityProcess(val ast.Argument) ProcessFunc {
	return func(c *CallStack) {
		c.ReturnValue = val
	}
}
