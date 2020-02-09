package parser

import (
	"errors"
	"fmt"
	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/lexer"
)

type (
	Parser struct {
		tokens []lexer.Lexem
		pos    int
	}
)

func ParseInteractive(tokens []lexer.Lexem) (*ast.Pipeline, error) {
	p := &Parser{
		tokens: tokens,
		pos:    -1,
	}
	if err := p.find(lexer.BeginCommand); err != nil {
		return nil, err
	}

	pline, err := parsePipeline(p)
	if err != nil {
		return nil, err
	}

	if err := p.find(lexer.EndCommand); err != nil {
		return nil, err
	}
	return pline, nil
}

func parsePipeline(p *Parser) (*ast.Pipeline, error) {
	pl := &ast.Pipeline{}
	for !p.eof() {
		cmd, err := parseOneCommand(p)
		if err != nil {
			return nil, err
		}
		pl.Items = append(pl.Items, cmd)
		if p.cur().Type == lexer.PipeConnector {
			return nil, errors.New("pipe not implemented")
		}
	}
	return pl, nil
}

func parseOneCommand(p *Parser) (*ast.Command, error) {
	cmd := &ast.Command{}
	if err := p.lookAhead(1, lexer.Identifier); err != nil {
		return nil, err
	}
	p.next()
	cmd.Identifier = p.cur()
	for p.next() {
		println("Cur: ", p.cur().Type.String())
		switch p.cur().Type {
		case lexer.EndCommand, lexer.PipeConnector:
			println("end command")
			p.prev()
			return cmd, nil
		case lexer.Identifier:
			cmd.Arguments = append(cmd.Arguments, &ast.Identifier{Name: p.cur()})
		case lexer.Number:
			cmd.Arguments = append(cmd.Arguments, &ast.Number{Value: p.cur()})
		case lexer.QuotedText:
			cmd.Arguments = append(cmd.Arguments, &ast.QuotedText{Text: p.cur()})
		}
	}
	return cmd, nil
}

func (p *Parser) find(tkType lexer.LexemType) error {
	for p.next() {
		if p.cur().Type == tkType {
			return nil
		}
	}
	return fmt.Errorf("expecting %v", tkType)
}

func (p *Parser) lookAhead(n int, tp lexer.LexemType) error {
	rewind := 0
	for ; n > 0; n-- {
		if p.next() {
			rewind++
		}
	}
	defer func() {
		for ; rewind > 0; rewind-- {
			p.prev()
		}
	}()
	if p.cur().Type != tp {
		return fmt.Errorf("expecting %v", tp)
	}
	return nil
}

func (p *Parser) prev() bool {
	if p.pos == -1 {
		return false
	}
	p.pos--
	return true
}

func (p *Parser) next() bool {
	if p.eof() {
		return false
	}
	p.pos++
	return true
}

func (p *Parser) cur() lexer.Lexem {
	if p.eof() {
		return lexer.Lexem{}
	}
	return p.tokens[p.pos]
}

func (p *Parser) eof() bool {
	return p.pos == len(p.tokens)-1
}
