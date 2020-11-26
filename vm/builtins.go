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
}
