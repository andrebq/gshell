package parser

import (
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
		{subject: "test single line commands", code: "{ echo hello world; }", fmt: "{ echo hello world }"},
		{subject: "test multiple command single line", code: "{ echo hello world; echo ola mundo; } ", fmt: "{\n\techo hello world\n\techo ola mundo\n}"},
		{subject: "test multiple command multiple lines", code: `{
	echo hello world
	echo ola mundo
}`},
	}
)

func TestParser(t *testing.T) {
	runTestCase := func(t *testing.T, tc testCase) {
		code := tc.code
		ast, err := Parse(code)
		if err != tc.expectedError {
			t.Errorf("Case: %v expected %v but got %v", tc.subject, tc.expectedError, err)
		}
		if ast.String() != tc.fmt {
			t.Errorf("Case: %v Should format to %q got %q", tc.subject, tc.fmt, ast.String())
		}
	}

	for _, tc := range allTests {
		runTestCase(t, tc)
	}

	t.Fail()
}
