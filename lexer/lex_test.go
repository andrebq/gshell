package lexer

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("..", "testdata", "one-liner", "simple.gshell"))
	if err != nil {
		t.Fatal(err)
	}
	lex, err := Lex(data)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(lex, []Lexem{
		Lexem{Identifier, "echo"},
		Lexem{QuotedText, "hello"},
		Lexem{Number, "1234"},
		Lexem{Identifier, "ola"},
	}) {
		t.Errorf("Should have 2 items. got: %v", lex)
	}
}
