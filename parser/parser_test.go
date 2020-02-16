package parser

import (
	. "github.com/andrebq/gshell/ast/astutil"
	"github.com/andrebq/gshell/lexer"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseInteractive(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("..", "testdata", "one-liner", "simple.gshell"))
	if err != nil {
		t.Fatal(err)
	}
	lex, err := lexer.Lex(data)
	if err != nil {
		t.Fatal(err)
	}

	pipeline, err := ParseInteractive(lex)
	if err != nil {
		t.Fatal(err)
	}

	expected := Pipeline(
		Command(
			Identifier("echo"), QuotedText("hello"), Number("1234"), Identifier("ola"), Number("234.34")))

	if !reflect.DeepEqual(pipeline, expected) {
		t.Errorf("Expecting %v got %v", expected, pipeline)
	}
}
