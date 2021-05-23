package mailbox

import (
	"context"
	"errors"
	"runtime"
	"sync"
)

type (
	Signal struct{}
	// Reader represents the read-only operations of a box,
	// note that by reading messages from a box the message
	// is removed, so while techinically this is a read operation,
	// it has side-effects
	Reader interface {
		// Depth returns the current size of this box
		// by the time the response is obtained the function
		// returns the actual value might be different.
		Depth() int
		Take() interface{}

		// Returns a channel that is closed when there are
		// data in the buffer
		WaitForData() <-chan Signal
	}
	// Box holds a queue where data can be buffered
	// for a given PID to consume
	//
	// While in Actor model the list of items for a given actor
	// is, logically, infinite. That is used as modeling simplification.
	//
	// Then underlying box allows the user to either:
	//
	// - Discard old items so writers are not blocked
	// - Discard new items so writer are not blocked
	// - Dicarsd the incoming item
	// - Block writers
	//
	// Readers always block while waiting for data
	Box interface {
		Reader

		// Reader returns the read-only version of this box
		Reader() Reader

		Push(interface{}) bool
		// Discard this mailbox which will prevent it from
		// reading any new input
		//
		// Items in the buffer are discarded too, this is
		// effectively cleaning up all the memory used by
		// this box
		Discard()

		// WaitForCapacity returns a channel which is closed
		// when the buffer acquires at least one slot.
		//
		// There is no guarantee that between the channel closed
		// and the next call to Push the buffer will remain
		// available.
		WaitForCapacity() <-chan Signal
	}
	box struct {
		lock sync.Mutex
		buf  []interface{}

		waitingForData []chan Signal
		waitingForSlot []chan Signal
	}

	readonly struct {
		wrapped *box
	}

	dropLogic byte
)

var (
	errMailboxNotRunning = errors.New("mailbox is not in a running state")
)

func NonBlocking(bufSize int) Box {
	return newBox(bufSize)
}

func newBox(bufSize int) Box {
	if bufSize <= 0 {
		bufSize = 1
	}
	b := &box{
		buf:            make([]interface{}, 0, bufSize),
		waitingForData: make([]chan Signal, 0, 1),
		waitingForSlot: make([]chan Signal, 0, runtime.GOMAXPROCS(-1)),
	}
	return b
}

func (b *box) WaitForData() <-chan Signal {
	ch := make(chan Signal)
	b.locked(func() {
		if len(b.buf) > 0 {
			close(ch)
			return
		}
		b.waitingForData = append(b.waitingForData, ch)
	})
	return ch
}

// Reader returns the reader interface of this box
func (b *box) Reader() Reader {
	return &readonly{wrapped: b}
}

func (b *box) Push(val interface{}) bool {
	if val == nil {
		return false
	}
	var written bool
	b.locked(func() {
		if len(b.buf) == cap(b.buf) {
			return
		}
		b.buf = append(b.buf, val)
		for idx, v := range b.waitingForData {
			close(v)
			b.waitingForData[idx] = nil
		}
		b.waitingForData = b.waitingForData[:0]
		written = true
	})
	return written
}

func (b *box) Take() interface{} {
	var ret interface{}
	b.locked(func() {
		if len(b.buf) == 0 {
			return
		}
		ret = b.buf[0]
		copy(b.buf[1:], b.buf[:len(b.buf)-1])
		b.buf[len(b.buf)-1] = nil
		b.buf = b.buf[:len(b.buf)-1]

		for i, v := range b.waitingForSlot {
			close(v)
			b.waitingForSlot[i] = nil
		}
		b.waitingForSlot = b.waitingForSlot[:0]
	})
	return ret
}

func (b *box) WaitForCapacity() <-chan Signal {
	ch := make(chan Signal)
	b.locked(func() {
		if len(b.buf) < cap(b.buf) {
			close(ch)
			return
		}
		b.waitingForSlot = append(b.waitingForSlot, ch)
	})
	return ch
}

func (b *box) Depth() int {
	return len(b.buf)
}

func (b *box) Discard() {
	b.locked(func() {
		b.buf = nil
	})
}

func (b *box) locked(fn func()) {
	b.lock.Lock()
	defer b.lock.Unlock()
	fn()
}

func (r *readonly) Depth() int {
	return r.wrapped.Depth()
}

func (r *readonly) Take() interface{} {
	return r.wrapped.Take()
}

func (r *readonly) WaitForData() <-chan Signal {
	return r.wrapped.WaitForData()
}

func BlockingPush(ctx context.Context, target Box, value interface{}) error {
	for {
		if target.Push(value) {
			return nil
		}
		select {
		case <-target.WaitForCapacity():
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func BlockingRead(ctx context.Context, target Reader) (interface{}, error) {
	for {
		value := target.Take()
		if value != nil {
			return value, nil
		}
		select {
		case <-target.WaitForData():
			continue
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
