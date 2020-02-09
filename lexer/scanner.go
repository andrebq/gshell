package lexer

import (
	"bytes"
	"errors"
	"unicode"
	"unicode/utf8"
)

type (
	scanner struct {
		code []rune
		pos  int
	}
)

func newScanner(c []byte) (*scanner, error) {
	if !utf8.Valid(c) {
		return nil, errors.New("not utf8")
	}
	s := &scanner{
		code: make([]rune, 0, utf8.RuneCount(c)),
		pos:  -1,
	}
	for len(c) > 0 {
		r, sz := utf8.DecodeRune(c)
		s.code = append(s.code, r)
		c = c[sz:]
	}
	return s, nil
}

func (s *scanner) eof() bool {
	return s.pos == len(s.code)-1
}

func (s *scanner) scanDoubleQuote() bool {
	if s.peek() == '"' {
		s.next()
		return true
	}
	return false
}

func (s *scanner) scanTextInQuotes() string {
	return s.scanWhile(func(r rune) bool { return r != '"' })
}

func (s *scanner) scanNumber() string {
	// TODO: handle float numbers here
	return s.scanWhile(unicode.IsDigit)
}

func (s *scanner) peek() rune {
	return s.code[s.pos+1]
}

func (s *scanner) next() bool {
	s.pos++
	return !s.eof()
}

func (s *scanner) scanWhile(filter func(r rune) bool) string {
	buf := bytes.Buffer{}
	for !s.eof() {
		if !filter(s.peek()) {
			break
		}
		buf.WriteRune(s.peek())
		s.next()
	}
	return buf.String()
}

func (s *scanner) scanNotWhitespace() string {
	return s.scanWhile(not(unicode.IsSpace))
}

func not(filter func(rune) bool) func(rune) bool {
	return func(r rune) bool {
		return !filter(r)
	}
}

func (s *scanner) discardWhitespace() {
	for !s.eof() && unicode.IsSpace(s.peek()) {
		s.next()
	}
}
