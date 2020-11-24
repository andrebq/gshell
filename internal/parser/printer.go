package parser

import "io"

type (
	Printer interface {
		WriteString(string)
		WriteIndent()
		WriteLineBreak()
		WriteArgSeparator()
		Indent()
		Unindent()
		Err() error
	}

	printer struct {
		nesting int
		out     io.Writer
		err     error
	}
)

func (p *printer) Indent() {
	p.nesting++
}

func (p *printer) Unindent() {
	p.nesting--
}

func (p *printer) WriteString(s string) {
	if p.err != nil {
		return
	}
	_, p.err = io.WriteString(p.out, s)
}

func (p *printer) WriteLineBreak() {
	p.WriteString("\n")
}

func (p *printer) WriteIndent() {
	for i := p.nesting; i > 0; i-- {
		p.WriteString("\t")
	}
}

func (p *printer) WriteArgSeparator() {
	p.WriteString(" ")
}

func (p *printer) Err() error {
	return p.err
}
