package ast

func (l *List) anchor() {}

// Head returns the first value of this list
func (l *List) Head() Argument {
	return l.val
}

// Tail returns the rest of the list except the first item
func (l *List) Tail() *List {
	return l.tail
}

// Nil returns true if the list is empty
func (l *List) Nil() bool {
	return l == nil
}

// Reverse returns the same items in reverse order
func (l *List) Reverse() *List {
	return doReverse(l, nil)
}

func doReverse(original *List, reversed *List) *List {
	// inspiration: https://stackoverflow.com/a/34437069
	if original.Nil() {
		return reversed
	}
	return doReverse(original.Tail(), Cons(original.Head(), reversed.Tail()))
}

// Fmt prints the cannonical representation of this list
func (l *List) Fmt(p Printer) {
	p.WriteString("[")
	l.fmt(p)
	p.WriteString("]")
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

// Cons returns a new list by appending value in front tail
// the tail isn't affected by this operation
func Cons(v Argument, tail *List) *List {
	return &List{val: v, tail: tail}
}

// NilList returns the empty list
func NilList() *List { return nil }
