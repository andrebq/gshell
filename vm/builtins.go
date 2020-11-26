package vm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andrebq/gshell/ast"
)

var (
	StdoutChannel = mustSym("stdout")
	StdinChannel  = mustSym("stdin")
	StderrChannel = mustSym("stderr")
)

func GShellPrintln(c *CallStack) {
	// TODO: this implementation of println is wrong, but I just want to see
	// if the function call works
	parts := make([]string, len(c.RawArgs))
	for i, a := range c.RawArgs {
		switch a := a.(type) {
		case ast.Symbol:
			parts[i] = a.Text()
		case ast.Number:
			parts[i] = strconv.FormatFloat(a.Float64(), 'f', -1, 64)
		default:
			parts[i] = fmt.Sprintf("%v", a)
		}
	}
	c.VM.PushValues(StdoutChannel, c.Context, strings.Join(parts, " ")+"\n")
}
