package vm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andrebq/gshell/ast"
)

func TestVM(t *testing.T) {
	vm := NewVM()
	// no need to consume all the tokens from stdout
	// as the buffer is large enough
	_, err := vm.Run("{ println hello world; }")
	if err != nil {
		t.Error(err)
	}

	out := vm.Stdout()
	var ev Event
	select {
	case ev = <-out:
		if !reflect.DeepEqual(ev.Main, fmt.Sprintf("hello world\n")) {
			t.Errorf("Unexpected value %#v in stdout channel", ev.Main)
		}
	default:
		t.Error("out channel should have at least one entry")
	}
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
	var ev Event
	select {
	case ev = <-out:
		if !reflect.DeepEqual(ev.Main, fmt.Sprintf("123\n")) {
			t.Errorf("Unexpected value %#v in stdout channel", ev.Main)
		}
	default:
		t.Error("out channel should have at least one entry")
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
	out := vm.Stdout()
	var ev Event
	select {
	case ev = <-out:
		if !reflect.DeepEqual(ev.Main, fmt.Sprintf("true\n")) {
			t.Errorf("Unexpected value %#v in stdout channel", ev.Main)
		}
	default:
		t.Error("out channel should have at least one entry")
	}
	select {
	case ev = <-out:
		if !reflect.DeepEqual(ev.Main, fmt.Sprintf("false\n")) {
			t.Errorf("Unexpected value %#v in stdout channel", ev.Main)
		}
	default:
		t.Error("out channel should have two entries")
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
	out := vm.Stdout()
	var ev Event
	select {
	case ev = <-out:
		if !reflect.DeepEqual(ev.Main, fmt.Sprintf("1\n")) {
			t.Errorf("Unexpected value %#v in stdout channel", ev.Main)
		}
	default:
		t.Error("out channel should have at least one entry")
	}
	select {
	case ev = <-out:
		if !reflect.DeepEqual(ev.Main, fmt.Sprintf("10\n")) {
			t.Errorf("Variable defined within a guard should be visible to the outer context. Got %v", ev.Main)
		}
	default:
		t.Error("out channel should have at least one entry")
	}
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
	actualEvents, err := extractAtLeastValues(vm.Stdout(), 5)
	if err != nil {
		t.Fatal(err)
	}
	expectedEvents := []Value{
		ast.NewText("1\n"),
		ast.NewText("2\n"),
		ast.NewText("3\n"),
		ast.NewText("4\n"),
		ast.NewText("5\n"),
	}
	if !reflect.DeepEqual(actualEvents, expectedEvents) {
		t.Errorf("Expecting %#v got %#v", expectedEvents, actualEvents)
	}
}

func extractAtLeastValues(in <-chan Event, n int) ([]Value, error) {
	buf := make([]Value, 0, n)
	for len(buf) < n {
		select {
		case ev := <-in:
			buf = append(buf, ev.Main)
		default:
			return buf, fmt.Errorf("Unable to read %v events from channel got %v", n, len(buf))
		}
	}
	return buf, nil
}
