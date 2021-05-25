package vm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andrebq/gshell/mailbox"
)

func TestVM(t *testing.T) {
	vm := NewVM()
	// no need to consume all the tokens from stdout
	// as the buffer is large enough
	_, err := vm.Run("{ println hello world; }")
	if err != nil {
		t.Error(err)
	}

	assertOutput(t, vm.Stdout(), []Value{
		"hello world\n",
	})
}

func TestSetVariable(t *testing.T) {
	vm := NewVM()
	// no need to consume all the tokens from stdout
	// as the buffer is large enough
	_, err := vm.Run("{ let $hello 123; println $hello; }")
	if err != nil {
		t.Error(err)
	}

	out := vm.Stdout()
	value := out.Take()
	if value == nil {
		t.Fatal("output should have at least one item")
	}
	if !reflect.DeepEqual(value, fmt.Sprintf("123\n")) {
		t.Errorf("Unexpected value %#v in stdout channel", value)
	}
}

func TestConditional(t *testing.T) {
	vm := NewVM()
	_, err := vm.Run(`{
		switch {
			case { true; } { println true; }
			else { println else; }
		}
		switch {
			case { false; } { println case-false; }
			else { println false; }
		}
}`)
	if err != nil {
		t.Error(err)
	}
	items, err := extractAtLeastValues(vm.Stdout(), 2)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(items, []Value{"true\n", "false\n"}) {
		t.Errorf("Unexpected values: %#v", items)
	}
}

func TestGuard(t *testing.T) {
	vm := NewVM()
	_, err := vm.Run(`{
		guard { false; } { println 0; }
		guard { true; } { println 1; }
		guard { true; } { let $a 10; }
		println $a
}`)
	if err != nil {
		t.Error(err)
	}
	assertOutput(t, vm.Stdout(), []Value{
		"1\n",
		"10\n",
	})
}

func TestSequenceLoop(t *testing.T) {
	vm := NewVM()
	_, err := vm.Run(`{
		loop i from 1 to 5 {
			println $i
		}
	}`)
	if err != nil {
		t.Fatal(err)
	}
	assertOutput(t, vm.Stdout(), []Value{
		"1\n",
		"2\n",
		"3\n",
		"4\n",
		"5\n",
	})
}
func TestFunctionDefinition(t *testing.T) {
	vm := NewVM()
	_, err := vm.Run(`{
		let $moduleVar 10
		let $otherModuleVar 10
		func print-a-and-b [$a $b] {
			let $localVar 20
			let $otherModuleVar 30
			println $a $b $moduleVar $localVar $otherModuleVar
		}
		print-a-and-b 10 [10 20]
		println $otherModuleVar
}`)
	if err != nil {
		t.Errorf("Unable to call-func with parameters: %v", err)
	} else {
		assertOutput(t, vm.Stdout(), []Value{
			"10 [ 10 20 ] 10 20 30\n",
			"30\n",
		})
	}
}

func assertOutput(t *testing.T, output mailbox.Reader, values []Value) []Value {
	items, err := extractAtLeastValues(output, len(values))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(items, values) {
		t.Errorf("Expecting %#v got %#v", values, items)
	}
	return items
}

func extractAtLeastValues(in mailbox.Reader, n int) ([]Value, error) {
	buf := make([]Value, 0, n)
	for len(buf) < n {
		v := in.Take()
		if v == nil {
			return nil, fmt.Errorf("Buffer was empty at %v elements", len(buf))
		}
		buf = append(buf, v)
	}
	return buf, nil
}

func TestFunctionsCannotDeclareOtherFunctions(t *testing.T) {
	vm := NewVM()
	_, err := vm.Run(`{
		func print-a-and-b [$a,$b] {
			func internal [] {
				true
			}
		}
		print-a-and-b 10 10
	}`)
	if err == nil {
		t.Fatal("Gshell should not support internal function definitions (yet!)")
	}
}
