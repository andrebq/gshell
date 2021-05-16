package mailbox

import (
	"context"
	"sync"
	"testing"
)

func benchmarSaturation(b *testing.B, bufSize, writersCount, readersCount int) {
	b.StopTimer()
	box := DiscardOldest(bufSize)
	var wg sync.WaitGroup
	wg.Add(writersCount)
	startWorkers := make(chan struct{})
	for i := 0; i < writersCount; i++ {
		go func(workerID int) {
			<-startWorkers
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				<-box.Push(workerID)
			}
		}(i)
	}

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < readersCount; i++ {
		go func() {
			select {
			case <-ctx.Done():
				return
			case <-box.Take():
			}
		}()
	}
	b.StartTimer()
	// broadcast to all workers that they should start
	// sending data
	close(startWorkers)
	wg.Wait()
	cancel()
}

func BenchmarkSaturation1x10x1(b *testing.B) {
	benchmarSaturation(b, 1, 10, 1)
}

func BenchmarkSaturation10x10x1(b *testing.B) {
	benchmarSaturation(b, 10, 10, 1)
}
func BenchmarkSaturation100x10x1(b *testing.B) {
	benchmarSaturation(b, 100, 10, 1)
}

func BenchmarkSaturation10x100x1(b *testing.B) {
	benchmarSaturation(b, 10, 100, 1)
}

func BenchmarkSaturation10x1000x1(b *testing.B) {
	benchmarSaturation(b, 10, 1000, 1)
}

func BenchmarkSaturation100x100x1(b *testing.B) {
	benchmarSaturation(b, 100, 100, 1)
}

func BenchmarkSaturation1000x1000x1(b *testing.B) {
	benchmarSaturation(b, 1000, 1000, 1)
}
