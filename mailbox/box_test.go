package mailbox

import (
	"reflect"
	"sync"
	"testing"
)

func pushAndWait(b Box, v interface{}, t *testing.T) error {
	err, open := <-b.Push(v)
	if !open {
		t.Errorf("box closed")
	}
	return err
}

func TestDiscardOlder(t *testing.T) {
	b := DiscardOldest(1)
	for i := 0; i < 10; i++ {
		if err := pushAndWait(b, i, t); err != nil {
			t.Error(err)
		}
	}
	if total := b.(*box).discardedItems; total != 9 {
		t.Errorf("should have discarded %v but got %v", 9, total)
	}
}

func TestBlocking(t *testing.T) {
	b := Blocking(2)
	if err := pushAndWait(b, 1, t); err != nil {
		t.Error(err)
	}
	if err := pushAndWait(b, 2, t); err != nil {
		t.Error(err)
	}
	select {
	case <-b.Push(1):
		t.Fatal("should not push since the buffer is supposed to be full")
	default:
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { <-b.Take(); wg.Done() }()
	go func() { <-b.Take(); wg.Done() }()
	wg.Wait()
}

func TestCanCloseInput(t *testing.T) {
	b := Blocking(2)
	<-b.Push(1)
	<-b.Push(2)
	b.CloseInput()
	if err := <-b.Push(3); err != errMailboxNotRunning {
		t.Fatal("After closing the input, calls to push should return an error!")
	}
	vals := []int{(<-b.Take()).(int), (<-b.Take()).(int)}
	if !reflect.DeepEqual(vals, []int{1, 2}) {
		t.Errorf("Should have two values %v got %v", []int{1, 2}, vals)
	}
}
