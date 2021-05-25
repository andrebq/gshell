package mailbox

import (
	"context"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func runIOOperations(ctx context.Context, t interface{ Fatal(...interface{}) }, timeout time.Duration, workers, readers, msgsPerWorker int) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	readerCtx, cancelReaders := context.WithCancel(ctx)
	defer cancel()
	defer cancelReaders()

	b := NonBlocking(workers)
	total := int64(workers * msgsPerWorker)
	{
		// readers
		for i := 0; i < readers; i++ {
			go func(i int) {
				for {
					if _, err := BlockingRead(readerCtx, b); err != nil {
						return
					}
					if atomic.AddInt64(&total, -1) <= 0 {
						cancelReaders()
						return
					}
				}
			}(i)
		}
	}
	{
		// writers
		for i := 0; i < workers; i++ {
			go func(i int) {
				for c := 0; c < msgsPerWorker; c++ {
					if err := BlockingPush(ctx, b, i); err != nil {
						return
					}
				}
			}(i)
		}
	}

	select {
	case <-readerCtx.Done():
		select {
		case <-ctx.Done():
			// the reader ctx was cancelled because the parent context
			// was canceled
			t.Fatal("Context cancelled because of parent", ctx.Err())
		default:
			cancel()
		}
	case <-ctx.Done():
		t.Fatal("Parent context canceled", ctx.Err())
	}
}

func benchIOOperations(ctx context.Context, t *testing.B, timeout time.Duration, workers, readers, msgsPerWorker int) {
	runIOOperations(ctx, t, timeout, workers, readers, t.N)
}

func TestMultiWriterSingleReader(t *testing.T) {
	runIOOperations(context.Background(), t, time.Second, 10, 1, 100)
}

func Benchmark100x1(t *testing.B) {
	benchIOOperations(context.Background(), t, time.Minute, 100, 1, 10)
}

func Benchmark1000x1(t *testing.B) {
	benchIOOperations(context.Background(), t, time.Minute, 1000, 1, 10)
}
func Benchmark100x4(t *testing.B) {
	benchIOOperations(context.Background(), t, time.Minute, 100, 4, 10)
}

func Benchmark1000x4(t *testing.B) {
	benchIOOperations(context.Background(), t, time.Minute, 1000, 4, 10)
}

func Benchmark1000xMaxProc(t *testing.B) {
	benchIOOperations(context.Background(), t, time.Minute, 1000, runtime.GOMAXPROCS(-1), 10)
}
