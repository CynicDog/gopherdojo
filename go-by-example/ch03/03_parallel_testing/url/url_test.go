package url

import (
	"sync"
	"testing"
	"time"
)

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestSequential(t *testing.T) {
}

func TestQuery(t *testing.T) {
	t.Parallel()
	t.Run("byName", func(t *testing.T) {
		t.Parallel()
		time.Sleep(2 * time.Second)
	})
	t.Run("byInventory", func(t *testing.T) {
		t.Parallel()
		time.Sleep(2 * time.Second)
	})
}

func incr(counter *int, mu *sync.Mutex) {
	mu.Lock()
	*counter++
	mu.Unlock()
}

func TestIncr(t *testing.T) {
	t.Parallel()

	t.Run("once", func(t *testing.T) {
		t.Parallel()
		var counter int
		var mu sync.Mutex
		incr(&counter, &mu)
		mu.Lock()
		defer mu.Unlock()
		if counter != 1 {
			t.Errorf("counter = %d, want 1", counter)
		}
	})

	t.Run("twice", func(t *testing.T) {
		t.Parallel()
		var counter int
		var mu sync.Mutex
		incr(&counter, &mu)
		incr(&counter, &mu)
		mu.Lock()
		defer mu.Unlock()
		if counter != 2 {
			t.Errorf("counter = %d, want 2", counter)
		}
	})
}
