package mailbox

import (
	"errors"
	"sync"
)

type (
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
		Push(interface{}) <-chan error
		Take() <-chan interface{}
		// Discard this mailbox which will prevent it from
		// reading any new input
		//
		// Items in the buffer are discarded too, this is
		// effectively cleaning up all the memory used by
		// this box
		Discard()

		// CloseInput disables any input for this box
		CloseInput()
	}
	box struct {
		lock      sync.Mutex
		dropLogic dropLogic

		buf []interface{}

		output      chan interface{}
		input       chan inputRequest
		inputClosed chan struct{}
		discard     chan struct{}

		discardedItems uint64
	}

	inputRequest struct {
		val      interface{}
		response chan error
	}

	dropLogic byte
)

const (
	doNotDrop = dropLogic(iota)
	dropIncoming
	dropHead
	dropTail
)

var (
	errMailboxNotRunning = errors.New("mailbox is not in a running state")
)

func Blocking(bufSize int) Box {
	return newBox(bufSize, doNotDrop)
}

func DiscardOldest(bufSize int) Box {
	return newBox(bufSize, dropHead)
}

func newBox(bufSize int, logic dropLogic) Box {
	if bufSize <= 0 {
		bufSize = 1
	}
	b := &box{
		dropLogic: logic,
		// this is closed when the box is no longer operating
		inputClosed: make(chan struct{}),
		discard:     make(chan struct{}),
		output:      make(chan interface{}),
		input:       make(chan inputRequest, 1000),
		buf:         make([]interface{}, 0, bufSize),
	}
	go b.run()
	return b
}

func (b *box) run() {
	defer b.doCloseInput()
	defer close(b.output)
	var nextInput inputRequest
	var nextOutput interface{}
	input := b.input
	output := b.output
	for {
		nextOutput = nil
		output = b.output
		// disable output if the buffer is empty
		// else take the head and use it as the
		// next output
		if len(b.buf) == 0 {
			output = nil
		} else {
			nextOutput = b.buf[0]
		}
		select {
		case <-b.inputClosed:
			input = nil
			b.input = nil
		default:
		}
		select {
		case output <- nextOutput:
			copy(b.buf[0:], b.buf[1:])
			b.buf = b.buf[:len(b.buf)-1]
			if input == nil {
				// input was disabled because the buffer was full
				// enable again since we just acquired a new slot
				input = b.input
			}
		case <-b.discard:
			return
		case nextInput = <-input:
			if len(b.buf) < cap(b.buf) {
				b.buf = append(b.buf, nextInput.val)
				nextInput.response <- nil
				close(nextInput.response)

				if b.dropLogic == doNotDrop &&
					len(b.buf) == cap(b.buf) {
					// if we cannot allow for input to be dropped
					// then disable the input since the buffer is full now
					input = nil
				}
			} else {
				b.discardedItems++
				switch b.dropLogic {
				case dropHead:
					copy(b.buf[0:], b.buf[1:])
					b.buf[len(b.buf)-1] = nextInput.val
					nextInput.response <- nil
					close(nextInput.response)
				case dropTail:
					b.buf[len(b.buf)-1] = nextInput.val
					nextInput.response <- nil
				case dropIncoming:
					nextInput.response <- nil
				case doNotDrop:
					panic("it must be impossible to get a buffer full with a doNotDrop logic. something went terribly wrong!")
				}
			}
		}
	}
}

func (b *box) Push(val interface{}) <-chan error {
	out := make(chan error, 1)
	ir := inputRequest{
		val:      val,
		response: out,
	}
	// try the fast non-blocking approach to avoid
	// too many goroutines
	if !b.nonBlockingPush(ir) {
		// go the memory intensive route
		// it might make sense to put a high limit here
		go b.blockingPush(ir)
	}
	return out
}

func (b *box) nonBlockingPush(ir inputRequest) bool {
	// try the fast route which does not trigger a new goroutine
	select {
	case <-b.inputClosed:
		ir.response <- errMailboxNotRunning
		close(ir.response)
		return true
	case <-b.discard:
		ir.response <- errMailboxNotRunning
		close(ir.response)
		return true
	case b.input <- ir:
		return true
	default:
		return false
	}
}

func (b *box) blockingPush(ir inputRequest) {
	select {
	case <-b.discard:
		ir.response <- errMailboxNotRunning
		close(ir.response)
	case <-b.inputClosed:
		ir.response <- errMailboxNotRunning
		close(ir.response)
	case b.input <- ir:
	}
}

func (b *box) Take() <-chan interface{} {
	return b.output
}

func (b *box) Discard() {
	b.locked(func() {
		select {
		case <-b.discard:
			return
		default:
			close(b.discard)
		}
	})
}

// Close input prevents the mailbox from receiving items
// but any output already in the buffer will remain available
// until the buffer is consumed or Discard is called.
func (b *box) CloseInput() {
	b.doCloseInput()
}

func (b *box) locked(fn func()) {
	b.lock.Lock()
	defer b.lock.Unlock()
	fn()
}

func (b *box) doCloseInput() {
	b.locked(func() {
		select {
		case <-b.inputClosed:
			return
		default:
			close(b.inputClosed)
		}
	})
}
