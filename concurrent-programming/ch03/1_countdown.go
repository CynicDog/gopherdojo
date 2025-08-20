package main

import (
	"fmt"
	"time"
)

func main() {
	// Because we pass its address to another goroutine, it escapes from the main
	// goroutine's stack to the heap
	count := 5

	// We pass a pointer so both goroutines share the same memory location.
	go countdown(&count)

	// This loop prints the current value of 'count' every 500ms.
	for count > 0 {
		time.Sleep(500 * time.Millisecond)

		// The CPU may have cached this value locally, but cache-coherency ensures
		// the main goroutine eventually sees updates from the countdown goroutine.
		fmt.Printf("Count is at %d\n", count)
	}
}

// countdown is a separate goroutine that decrements 'seconds' every second.
func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)

		// Decrement the shared variable.
		// This modifies the value in heap memory, visible to all goroutines.
		*seconds -= 1
	}
}
