package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		// Writers (append) must take the full Lock (exclusive access).
		// This blocks both readers and writers until done.
		mutex.Lock()
		*matchEvents = append(*matchEvents,
			"Match event "+strconv.Itoa(i))
		mutex.Unlock()

		time.Sleep(200 * time.Millisecond)
		fmt.Println("Appended match event")
	}
}

func clientHandler(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 100; i++ {
		// Readers only need a shared read lock.
		// With RWMutex, many goroutines can hold RLock() simultaneously
		// as long as no one holds a write Lock().
		mutex.RLock()
		allEvents := copyAllEvents(mEvents)
		mutex.RUnlock()

		// This means thousands of clientHandlers can read concurrently,
		// instead of being forced into a single-file line like with sync.Mutex.
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	// Just copying the slice safely after acquiring a lock
	allEvents := make([]string, 0, len(*matchEvents))
	for _, e := range *matchEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}

func main() {
	mutex := sync.RWMutex{}
	var matchEvents = make([]string, 0, 10000)

	// Prepopulate with some events
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}

	// Writer goroutine (exclusive writes)
	go matchRecorder(&matchEvents, &mutex)

	start := time.Now()

	// Thousands of readers can now run concurrently,
	// which scales much better than with sync.Mutex.
	for j := 0; j < 5000; j++ {
		go clientHandler(&matchEvents, &mutex, start)
	}

	time.Sleep(100 * time.Second)
}
