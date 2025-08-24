package main

import (
	"fmt"
	"sync"
	"time"
)

// ReadWriteMutex A simplified implementation of a Read-Write mutex.
//
// - multiple readers can enter concurrently
// - writers get exclusive access (no readers or writers at the same time)
//
// Internally it uses:
//   - readersLock: a short-term mutex just to protect the readersCounter
//   - readersCounter: how many readers are active right now
//   - globalLock: the "big" lock that enforces exclusion between
//     readers (as a group) and writers
type ReadWriteMutex struct {
	readersCounter int        // number of active readers
	readersLock    sync.Mutex // protects readersCounter
	globalLock     sync.Mutex // blocks writers and readers group against each other
}

func (rw *ReadWriteMutex) ReadLock() {
	// Step 1: readers must update the readersCounter safely,
	// so they acquire readersLock.
	rw.readersLock.Lock()

	// Step 2: increment active readers
	rw.readersCounter++

	// Step 3: if this is the FIRST reader,
	// it grabs the globalLock. That prevents writers
	// from entering while at least one reader is active.
	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}

	// Step 4: release readersLock quickly, so that
	// other readers can also increment the counter.
	// Note: this is the serialization point — all readers
	// pass through here briefly, but then they run concurrently.
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	// Step 1: acquire readersLock to safely decrement counter
	rw.readersLock.Lock()

	// Step 2: decrement active readers
	rw.readersCounter--

	// Step 3: if this was the LAST reader leaving,
	// release the globalLock so a waiting writer can proceed.
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock()
	}

	// Step 4: release readersLock so other readers/unlocks can proceed
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	// Writers take the globalLock directly.
	// This excludes both new readers and other writers.
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	// Writers release the globalLock.
	rw.globalLock.Unlock()
}

func main() {
	rwMutex := ReadWriteMutex{}

	// Start 10 readers. They should be able to run concurrently.
	for i := 0; i < 10; i++ {
		go func(id int) {
			rwMutex.ReadLock()
			fmt.Printf("Reader %d started\n", id)
			time.Sleep(5 * time.Second) // simulate long read
			fmt.Printf("Reader %d done\n", id)
			rwMutex.ReadUnlock()
		}(i)
	}

	// Give readers some time to start before trying to write
	time.Sleep(1 * time.Second)

	// Try to write. This will block until all readers finish,
	// because the first reader still holds the globalLock.
	fmt.Println("Writer waiting...")
	rwMutex.WriteLock()
	fmt.Println("Writer acquired lock, writing...")
	rwMutex.WriteUnlock()
	fmt.Println("Writer finished")
}
