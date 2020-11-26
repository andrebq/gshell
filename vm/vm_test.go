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
	err := vm.Run("{ println hello world; }")
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
