package shell

import (
	"testing"
)

func TestSimpleEcho(t *testing.T) {
	vm := NewVM()
	err := vm.ParseInteractive("simple-echo", "echo abc 1234")
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Run("simple-echo")
	if err != nil {
		t.Fatal(err)
	}
}
