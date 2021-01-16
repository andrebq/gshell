package ast

import "github.com/andrebq/gshell/internal/pdata"

func (l *List) anchor() {}

// Head returns the first value of this list
func (l *List) Head() Argument {
	if l.Nil() {
		return nil
	}
	return l.data.Get(0).(Argument)
}

// Tail returns the rest of the list except the first item
func (l *List) Tail() *List {
	if l.Nil() {
		return NilList()
	}
	return &List{data: l.data.Slice(1, l.data.Len())}
}

// Nil returns true if the list is empty
func (l *List) Nil() bool {
	return l.data.Len() == 0
}

// Reverse returns the same items in reverse order
func (l *List) Reverse() *List {
	slc := l.ToSlice(make([]Argument, 0, l.data.Len()))
	for i := 0; i < len(slc)/2; i++ {
		slc[i], slc[len(slc)-i-1] = slc[len(slc)-i-1], slc[i]
	}
	return NilList().Append(slc...)
}

// Append returns a new list with a copy of this list
// followed by the items
func (l *List) Append(values ...Argument) *List {
	l2 := &List{}
	aux := make([]pdata.Any, len(values))
	for i, v := range values {
		aux[i] = v
	}
	l2.data = l.data.Append(aux...)
	return l2
}

// Fmt prints the cannonical representation of this list
func (l *List) Fmt(p Printer) {
	p.WriteString("[")
	l.fmt(p)
	p.WriteString("]")
}

// ToSlice copies all the items from this list into
// a slice (which might be nil)
func (l *List) ToSlice(dest []Argument) []Argument {
	l.data.Range(func(a pdata.Any) bool {
		dest = append(dest, a.(Argument))
		return true
	})
	return dest
}

// fmt is just a helper to print all items while
// not printing the list head
func (l *List) fmt(p Printer) {
	p.WriteArgSeparator()
	if l.Nil() {
		return
	}
	l.Head().Fmt(p)
	l.Tail().fmt(p)
}

var nilList = &List{
	data: pdata.NewArgumentListSlice(),
}

// NilList returns the empty list
func NilList() *List { return nilList }
