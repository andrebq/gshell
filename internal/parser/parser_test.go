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
		{subject: "test empty command blocks", code: "{}", fmt: "{}"},
		{subject: "test last command does not require terminator", code: "{ hello world }", fmt: "{ hello world }"},
		{subject: "test single line commands", code: "{ echo hello world; }", fmt: "{ echo hello world }"},
		{subject: "test multiple command single line", code: "{ echo hello world; echo ola mundo } ", fmt: "{\n\techo hello world\n\techo ola mundo\n}"},
		{subject: "test multiple command multiple lines", code: `{
	echo hello world
	echo ola mundo
}`, fmt: "{\n\techo hello world\n\techo ola mundo\n}"},
		{subject: "test numeric arguments",
			code: "{echo 123 world123 w123h; }",
			fmt:  "{ echo 123 world123 w123h }"},
		{subject: "symbols have separators",
			code: "{ echo-1! should.be.valid }",
			fmt:  "{ echo-1! should.be.valid }"},
		{subject: "can use variables as arguments",
			code: "{println $variable }",
			fmt:  "{ println $variable }"},
		{subject: "can use blocks as arguments",
			code: "{ if true { println false; } else { println true; } }",
			fmt:  "{ if true { println false } else { println true } }"},
		{subject: "can parse lists as argument",
			code: `{ echo [ 123  $abc identifier ] }`,
			fmt:  `{ echo [ 123 $abc identifier ] }`},
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
			t.Errorf("Case: %v Formatted error caused a different error: %v", tc.subject, err)
		}
		if !reflect.DeepEqual(ast, secondAst) {
			t.Errorf("Case: %v Generated AST's do not match", tc.subject)
		}
	}

	for _, tc := range allTests {
		runTestCase(t, tc)
	}
}
