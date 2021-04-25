package ast

import (
	"reflect"
	"testing"
)

var _ Argument = NilList()

func TestListReverse(t *testing.T) {
	lst := NilList().Append(NewNumber(10), NewNumber(20), NewNumber(30))
	expected := []Argument{
		NewNumber(30),
		NewNumber(20),
		NewNumber(10),
	}
	actual := lst.Reverse().ToSlice(nil)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Reverse should generate %v got %v", expected, actual)
	}
}
