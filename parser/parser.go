package parser

import (
	"errors"
	"fmt"

	"github.com/andrebq/gshell/ast"
	"github.com/andrebq/gshell/ast/astutil"
	"github.com/andrebq/gshell/lexer"
)

type (
	Parser struct {
		tokens []lexer.Lexem
		pos    int
	}
)

func ParseProgram(tokens []lexer.Lexem) (*ast.Program, error) {
	p := &Parser{
		tokens: tokens,
		pos:    -1,
	}
	if err := p.find(lexer.OpenProgram); err != nil {
		return nil, err
	}
	p.prev()
	return parseProgram(p)
}

func ParseInteractive(tokens []lexer.Lexem) (*ast.Pipeline, error) {
	p := &Parser{
		tokens: tokens,
		pos:    -1,
	}
	if err := p.find(lexer.Identifier); err != nil {
		return nil, err
	}
	p.prev()

	pline, err := parsePipeline(p)
	if err != nil {
		return nil, err
	}
	return pline, nil
}

func parseProgram(p *Parser) (*ast.Program, error) {
	err := p.find(lexer.OpenProgram)
	if err != nil {
		return nil, err
	}

	prog := &ast.Program{}
	for !p.eof() && p.cur().Type != lexer.CloseProgram {
		if p.cur().Type == lexer.RawComment {
			// we should handle comments, but for now, lets ignore it
			p.next()
			continue
		}
		pipeline, err := parsePipeline(p)
		if err != nil {
			return nil, err
		}
		prog.Instructions = append(prog.Instructions, pipeline)
	}
	if p.cur().Type != lexer.CloseProgram {
		return nil, errors.New("missing CloseProgram")
	}
	return prog, nil
}

func parsePipeline(p *Parser) (*ast.Pipeline, error) {
	pl := &ast.Pipeline{}
	for !p.eof() {
		err := p.find(lexer.Identifier)
		if err != nil {
			return nil, err
		}
		p.prev()
		cmd, err := parseOneCommand(p)
		if err != nil {
			return nil, err
		}
		pl.Items = append(pl.Items, cmd)
		if p.cur().Type == lexer.Terminator {
			break
		} else if p.cur().Type == lexer.CloseProgram {
			break
		} else if p.cur().Type == lexer.PipeConnector {
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
	cmd.Identifier = astutil.Identifier(p.cur().Value)
	var commandComplete bool
	for p.next() && !commandComplete {
		switch p.cur().Type {
		case lexer.PipeConnector:
			p.prev()
			commandComplete = true
		case lexer.Terminator:
			commandComplete = true
		case lexer.CloseProgram:
			return nil, errors.New("every command should have a terminator")
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
	return !p.eof()
}

func (p *Parser) cur() lexer.Lexem {
	if p.pos == len(p.tokens) {
		return lexer.Lexem{}
	}
	return p.tokens[p.pos]
}

func (p *Parser) eof() bool {
	return p.pos == len(p.tokens)
}
