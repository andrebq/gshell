package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andrebq/gshell/ast"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type (
	stack []interface{}

	astBuilder struct {
		*BaseGShellListener
		ast *ast.Ast

		err error

		stack stack
	}
)

func newAstBuilder() *astBuilder {
	return &astBuilder{
		BaseGShellListener: &BaseGShellListener{},
		ast:                ast.New(),
	}
}

func Parse(code string) (*ast.Ast, error) {
	lexer := NewGShellLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGShellParser(stream)
	errorsFound := &errorsListener{
		ErrorListener: antlr.NewDefaultErrorListener(),
	}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorsFound)
	astBuilder := newAstBuilder()
	antlr.ParseTreeWalkerDefault.Walk(astBuilder, parser.Start())

	if errorsFound.err != nil {
		return nil, errorsFound.err
	}
	// TODO: this is probably a tech-debt as I assume it should
	// be possible to inject an error from the Listener
	//
	// but I'm too lazy to find the solution right now
	if astBuilder.err != nil {
		return nil, astBuilder.err
	}
	return astBuilder.ast, nil
}

func (ab *astBuilder) ExitStart(c *StartContext) {
	ab.ast.SetRoot(ab.stack.pop().(*ast.Script))
}

func (ab *astBuilder) EnterScript(c *ScriptContext) {
	ab.stack.push(ast.NewScript())
}

func (ab *astBuilder) ExitScript(c *ScriptContext) {
	steps := new(stack)
	for !isScript(ab.stack) {
		steps.push(ab.stack.pop().(*ast.Cmd))
	}

	sc := ab.stack.pop().(*ast.Script)
	for !steps.empty() {
		sc.AddCommand(steps.pop().(*ast.Cmd))
	}
	ab.stack.push(sc)
}

func (ab *astBuilder) EnterScriptArgument(c *ScriptArgumentContext) {
	ab.stack.push(ast.NewScript())
}

func (ab *astBuilder) ExitScriptArgument(c *ScriptArgumentContext) {
	steps := new(stack)
	for !isScript(ab.stack) {
		steps.push(ab.stack.pop().(*ast.Cmd))
	}
	sc := ab.stack.pop().(*ast.Script)
	for !steps.empty() {
		sc.AddCommand(steps.pop().(*ast.Cmd))
	}
	ab.stack.push(sc)
}

func (ab *astBuilder) EnterSingleCommand(c *SingleCommandContext) {
	ab.stack.push(&ast.Cmd{})
}

func (ab *astBuilder) ExitSingleCommand(c *SingleCommandContext) {
	argst := new(stack)
	for !isCmd(ab.stack) {
		argst.push(ab.stack.pop())
	}

	cmd := ab.stack.pop().(*ast.Cmd)

	for !argst.empty() {
		if cmd.Command() == (ast.Symbol{}) {
			cmd.SetCommand(argst.pop().(ast.Symbol))
			continue
		}
		cmd.AddArgument(argst.pop().(ast.Argument))
	}

	ab.stack.push(cmd)
}

func (ab *astBuilder) ExitCommandName(c *CommandNameContext) {
	s, err := ast.NewSymbol(c.GetText())
	if err != nil {
		ab.err = fmt.Errorf("string %q could not be cast to ast.Symbol. cause: %v", c.GetText(), err)
		s, _ = ast.NewSymbol("error-invalid-symbol")
	}
	ab.stack.push(s)
}

func (ab *astBuilder) ExitNumericArgument(c *NumericArgumentContext) {
	number, err := strconv.ParseFloat(c.GetText(), 63)
	if err != nil {
		ab.stack.push(ast.NewText(c.GetText()))
	} else {
		ab.stack.push(ast.NewNumber(number))
	}
}

func (ab *astBuilder) ExitNamedArgument(c *NamedArgumentContext) {
	s, err := ast.NewSymbol(c.GetText())
	if err != nil {
		ab.err = fmt.Errorf("string %q could not be cast to ast.Symbol. cause: %v", c.GetText(), err)
		s, _ = ast.NewSymbol("error-invalid-symbol")
	}
	ab.stack.push(s)
}

func (ab *astBuilder) ExitVariableArgument(c *VariableArgumentContext) {
	v, err := ast.NewVarString(c.GetText()[1:])
	if err != nil {
		ab.err = fmt.Errorf("string %q could not be cast to ast.Symbol. cause: %v", c.GetText(), err)
		v, _ = ast.NewVarString("error-invalid-variable")
	}
	ab.stack.push(v)
}

type listStartMarker struct{}

func (ab *astBuilder) EnterListArgument(c *ListArgumentContext) {
	ab.stack.push(listStartMarker{})
}

func (ab *astBuilder) ExitListArgument(c *ListArgumentContext) {
	lst := ast.NilList()
	for !ab.stack.empty() {
		v := ab.stack.pop()
		if _, isListStart := v.(listStartMarker); isListStart {
			break
		}
		lst = lst.Append(v.(ast.Argument))
	}
	ab.stack.push(lst.Reverse())
}

func (ab *astBuilder) ExitTextArgument(c *TextArgumentContext) {
	txt := c.GetText()
	if strings.HasPrefix(txt, `"""`) {
		txt = txt[3 : len(txt)-3]
	} else {
		txt = txt[1 : len(txt)-1]
	}
	ab.stack.push(ast.NewText(txt))
}

func (s *stack) push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) pop() interface{} {
	v := (*s)[len(*s)-1]
	// ensure gc cannot reach the old value
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) peek() interface{} {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool { return len(*s) == 0 }
func (s *stack) depth() int  { return len(*s) }

func isCmd(st stack) bool {
	_, ok := st.peek().(*ast.Cmd)
	return ok
}

func isScript(st stack) bool {
	_, ok := st.peek().(*ast.Script)
	return ok
}
