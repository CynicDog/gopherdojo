package main

import (
	"fmt"
	"runtime"
	"time"
)

// scheduledStingy repeatedly adds 10 to the shared money variable.
// WARNING: This operation (*money += 10) is NOT atomic.
// Even if we call runtime.Gosched(), the goroutine may be
// preempted between the read-modify-write sequence, leading
// to a race condition when another goroutine writes concurrently.
func scheduledStingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
		runtime.Gosched() // voluntarily yields, but does NOT prevent races
	}
	fmt.Println("Stingy Done")
}

// scheduledSpendy repeatedly subtracts 10 from the shared money variable.
// WARNING: Same issue as above. Concurrent modifications without
// synchronization may overwrite each other's changes.
func scheduledSpendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10

		// Voluntarily yields execution to another goroutine on the same thread
		// but this does NOT enforce atomicity or memory synchronization
		// so another goroutine (possibly on another CPU core) can still read/write money concurrently
		runtime.Gosched()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	// Two goroutines concurrently modify money without synchronization.
	// Result is unpredictable due to race conditions.
	go scheduledStingy(&money)
	go scheduledSpendy(&money)

	time.Sleep(2 * time.Second)
	println("Money in the bank account: ", money)
}
