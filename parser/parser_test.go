package parser

import (
	"github.com/andrebq/gshell/ast"
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
	expected := &ast.Pipeline{}
	if !reflect.DeepEqual(pipeline, expected) {
		t.Errorf("Expecting %v got %v", expected, pipeline)
	}
}
