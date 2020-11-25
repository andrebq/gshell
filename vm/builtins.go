package vm

var (
	StdoutChannel = mustSym("stdout")
	StdinChannel  = mustSym("stdin")
	StderrChannel = mustSym("stderr")
)

func GShellPrintln(c *CallStack) {
	c.VM.PushValues(StdoutChannel, "Call made with", len(c.RawArgs))
}
