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

func newScanner(c []byte, injectNewline bool) (*scanner, error) {
	if !utf8.Valid(c) {
		return nil, errors.New("not utf8")
	}
	s := &scanner{
		code: make([]rune, 0, utf8.RuneCount(c)+1),
		pos:  -1,
	}
	if bytes.HasSuffix(c, []byte("\n")) && injectNewline {
		// no need to inject the newline because the line already contains a newline
		injectNewline = false
	}
	for len(c) > 0 {
		r, sz := utf8.DecodeRune(c)
		s.code = append(s.code, r)
		c = c[sz:]
	}
	if injectNewline {
		s.code = append(s.code, '\n')
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
	buf := bytes.Buffer{}
	foundDot := false
	for !s.eof() {
		if unicode.IsDigit(s.peek()) {
			buf.WriteRune(s.read())
		} else if s.peek() == '.' && !foundDot {
			buf.WriteRune(s.read())
		} else {
			break
		}
	}
	return buf.String()
}

func (s *scanner) unread() {
	s.pos--
}

func (s *scanner) read() rune {
	p := s.peek()
	s.next()
	return p
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
