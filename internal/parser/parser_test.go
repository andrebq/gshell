package parser

import (
	"reflect"
	"testing"
)

type (
	testCase struct {
		subject       string
		code          string
		fmt           string
		expectedError error
	}
)

var (
	allTests = []testCase{
		{subject: "test single line commands", code: "{ echo hello world; }", fmt: "{ echo hello world; }"},
		{subject: "test multiple command single line", code: "{ echo hello world; echo ola mundo; } ", fmt: "{\n\techo hello world\n\techo ola mundo\n}"},
		{subject: "test multiple command multiple lines", code: `{
	echo hello world
	echo ola mundo
}`, fmt: "{\n\techo hello world\n\techo ola mundo\n}"},
		{subject: "test numeric arguments", code: "{echo 123 world123 w123h; }",
			fmt: "{ echo 123 world123 w123h; }"},
		{subject: "symbols have separators", code: "{ echo-1! should.be.valid; }",
			fmt: "{ echo-1! should.be.valid; }"},
	}
)

func TestParser(t *testing.T) {
	runTestCase := func(t *testing.T, tc testCase) {
		code := tc.code
		ast, err := Parse(code)
		if err != tc.expectedError {
			t.Errorf("Case: %v expected %v but got %v", tc.subject, tc.expectedError, err)
			return
		}
		fmtAst := ast.String()
		if fmtAst != tc.fmt {
			t.Errorf("Case: %v Should format to %q got %q", tc.subject, tc.fmt, fmtAst)
		}
		secondAst, err := Parse(fmtAst)
		if err != tc.expectedError {
			t.Errorf("Formatted error caused a different error: %v", err)
		}
		if !reflect.DeepEqual(ast, secondAst) {
			t.Errorf("Generated AST's do not match")
		}
	}

	for _, tc := range allTests {
		runTestCase(t, tc)
	}
}
