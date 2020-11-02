package parser

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	_ "github.com/andrebq/gshell/ast/astutil"
	"github.com/andrebq/gshell/lexer"
)

func TestParseInteractive(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("..", "testdata", "one-liner", "simple.gshell"))
	if err != nil {
		t.Fatal(err)
	}
	tokens, err := lexer.Lex(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Tokens: %v", tokens)
	t.Logf("Semantics have changed")
	t.FailNow()

	// pipeline, err := ParseInteractive(lex)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// expected := Pipeline(Command(
	// 	Identifier("echo"), QuotedText("hello"), Number("1234"), Identifier("ola"), Number("234.34")))

	// if !reflect.DeepEqual(pipeline, expected) {
	// 	t.Errorf("Expecting %v got %v", expected, pipeline)
	// }
}

func TestParseProgram(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("..", "testdata", "programs", "basic.gshell"))
	if err != nil {
		t.Fatal(err)
	}
	tokens, err := lexer.Lex(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Tokens: %v", tokens)
	t.Logf("Semantics have changed")
	t.FailNow()

	// program, err := ParseProgram(tokens)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// expected := Program(
	// 	Pipeline(Command(
	// 		Identifier("echo"), QuotedText("hello"), Number("123"))))

	// if !reflect.DeepEqual(program, expected) {
	// 	t.Errorf("Expecting %v got %v", expected, program)
	// }
}
