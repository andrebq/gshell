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
	expected := []Lexem{
		Lexem{Identifier, "echo"},
		Lexem{QuotedText, "hello"},
		Lexem{Number, "1234"},
		Lexem{Identifier, "ola"},
		Lexem{Number, "234.34"},
		Lexem{Terminator, ""},
	}
	if !reflect.DeepEqual(lex, expected) {
		t.Errorf("Should be %v got: %v", expected, lex)
	}
}

func TestPipeline(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("..", "testdata", "one-liner", "pipeline.gshell"))
	if err != nil {
		t.Fatal(err)
	}
	lex, err := Lex(data)
	if err != nil {
		t.Fatal(err)
	}
	expected := []Lexem{
		Lexem{Identifier, "echo"},
		Lexem{QuotedText, "hello"},
		Lexem{PipeConnector, "|"},
		Lexem{Identifier, "cat"},
		Lexem{Identifier, "output.txt"},
		Lexem{Terminator, ""},
	}
	if !reflect.DeepEqual(lex, expected) {
		t.Errorf("Should be %v got %v", expected, lex)
	}
}
