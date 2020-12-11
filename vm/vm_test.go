package vm

import (
	"fmt"
	"reflect"
	"testing"
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
