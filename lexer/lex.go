package lexer

import (
	"errors"
	"fmt"
	"unicode"
)

type (
	Lexem struct {
		Type  LexemType
		Value string
	}

	lexemOrErr struct {
		l Lexem
		e error
	}

	LexemType byte

	stateFn func(chan<- lexemOrErr, *scanner) stateFn
)

const (
	Undefined = LexemType(iota)
	Identifier
	OpenProgram
	CloseProgram
	Whitespace
	QuotedText
	Number
)

func (l Lexem) String() string {
	return fmt.Sprintf("%d:%v", l.Type, l.Value)
}

func Lex(code []byte) ([]Lexem, error) {
	sc, err := newScanner(code)
	if err != nil {
		return nil, err
	}
	_ = sc
	out := make(chan lexemOrErr)
	go lexer(out, sc)

	var lexems []Lexem
	for r := range out {
		if r.e != nil {
			return nil, r.e
		}
		lexems = append(lexems, r.l)
	}
	return lexems, nil
}

func lexer(out chan<- lexemOrErr, sc *scanner) {
	defer close(out)
	initialState := lexStart
	for curSt := initialState; !sc.eof() && curSt != nil; curSt = curSt(out, sc) {
	}
}

func lexStart(out chan<- lexemOrErr, sc *scanner) stateFn {
	sc.discardWhitespace()
	return lexNext
}

func lexIdentifier(out chan<- lexemOrErr, sc *scanner) stateFn {
	id := sc.scanNotWhitespace()
	if len(id) > 0 {
		out <- lexemOrErr{l: Lexem{Type: Identifier, Value: id}}
	}
	return lexNext
}

func lexDoubleQuoteText(out chan<- lexemOrErr, sc *scanner) stateFn {
	if !sc.scanDoubleQuote() {
		out <- lexemOrErr{e: errors.New("expecting double quotes")}
		return nil
	}

	txt := sc.scanTextInQuotes()

	if !sc.scanDoubleQuote() {
		out <- lexemOrErr{e: errors.New("expecting double quotes")}
		return nil
	}

	out <- lexemOrErr{l: Lexem{Type: QuotedText, Value: txt}}

	return lexNext
}

func lexNumber(out chan<- lexemOrErr, sc *scanner) stateFn {
	number := sc.scanNumber()
	if len(number) == 0 {
		out <- lexemOrErr{e: errors.New("expecting a number")}
		return nil
	}
	out <- lexemOrErr{l: Lexem{Type: Number, Value: number}}
	return lexNext
}

func lexNext(out chan<- lexemOrErr, sc *scanner) stateFn {
	r := sc.peek()
	if r == '{' {
		out <- lexemOrErr{l: Lexem{Type: OpenProgram, Value: "{"}}
	} else if r == '}' {
		out <- lexemOrErr{l: Lexem{Type: CloseProgram, Value: "}"}}
	} else if unicode.IsLetter(r) {
		return lexIdentifier
	} else if unicode.IsSpace(r) {
		sc.discardWhitespace()
		return lexNext
	} else if unicode.IsDigit(r) {
		return lexNumber
	} else if r == '"' {
		return lexDoubleQuoteText
	} else {
		out <- lexemOrErr{e: fmt.Errorf("unexpected character %v", string(sc.peek()))}
	}
	sc.next()
	return lexNext
}
