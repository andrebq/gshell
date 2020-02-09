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
	BeginCommand
	EndCommand
	Identifier
	OpenProgram
	CloseProgram
	Whitespace
	QuotedText
	Number
	OrOperator
	PipeConnector
	RawComment
)

func (l Lexem) String() string {
	return fmt.Sprintf("%d:%v", l.Type, l.Value)
}

func Lex(code []byte) ([]Lexem, error) {
	sc, err := newScanner(code, true)
	if err != nil {
		return nil, err
	}
	_ = sc
	out := make(chan lexemOrErr)
	go lexer(out, sc)
	filtered := make(chan lexemOrErr)
	go filter(filtered, out)

	var lexems []Lexem
	for r := range filtered {
		if r.e != nil {
			return nil, r.e
		}
		lexems = append(lexems, r.l)
	}
	return lexems, nil
}

func filter(out chan<- lexemOrErr, in <-chan lexemOrErr) {
	first := true
	endCommand := false
	defer close(out)
	for l := range in {
		if l.e != nil {
			out <- l
			continue
		}
		if l.l.Type == RawComment {
			out <- l
			continue
		}
		if first {
			first = false
			out <- lexemOrErr{l: Lexem{Type: BeginCommand, Value: ""}}
			out <- l
			continue
		}
		if l.l.Type == EndCommand && endCommand {
			// no need to duplicate end-command discard the current item
			continue
		} else if l.l.Type == EndCommand {
			// this is the first EndCommand token, let it pass and toggle endCommand so we can filter out the other ones
			endCommand = true
			out <- l
			continue
		}
		endCommand = false
		out <- l
	}
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

func lexPipeline(out chan<- lexemOrErr, sc *scanner) stateFn {
	first := sc.peek()
	if first != '|' {
		out <- lexemOrErr{e: errors.New("expecting a |")}
		return nil
	}
	sc.next()
	if sc.peek() == '|' {
		sc.next()
		out <- lexemOrErr{l: Lexem{Type: OrOperator, Value: "||"}}
		return lexNext
	}
	out <- lexemOrErr{l: Lexem{Type: PipeConnector, Value: "|"}}
	return lexNext
}

func lexNext(out chan<- lexemOrErr, sc *scanner) stateFn {
	r := sc.read()
	if r == '{' {
		out <- lexemOrErr{l: Lexem{Type: OpenProgram, Value: "{"}}
	} else if r == '}' {
		out <- lexemOrErr{l: Lexem{Type: CloseProgram, Value: "}"}}
	} else if unicode.IsLetter(r) {
		sc.unread()
		return lexIdentifier
	} else if unicode.IsSpace(r) && r != '\n' {
		sc.unread()
		sc.discardWhitespace()
		return lexNext
	} else if unicode.IsDigit(r) {
		sc.unread()
		return lexNumber
	} else if r == '"' {
		sc.unread()
		return lexDoubleQuoteText
	} else if r == '|' {
		sc.unread()
		return lexPipeline
	} else if r == '\n' {
		out <- lexemOrErr{l: Lexem{Type: EndCommand, Value: ""}}
	} else {
		sc.unread()
		out <- lexemOrErr{e: fmt.Errorf("unexpected character %v", string(sc.peek()))}
		sc.next()
	}
	return lexNext
}
