package main

import (
	"fmt"
	"sync"
	"time"
)

// stingy repeatedly adds 10 to money while holding the mutex.
// The mutex ensures that only one goroutine at a time can
// access the shared variable, preventing race conditions.
func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()   // enter critical section
		*money += 10   // safe update
		mutex.Unlock() // leave critical section
	}
	fmt.Println("Stingy Done")
}

// spendy repeatedly subtracts 10 from money while holding the mutex.
// Same as above: ensures mutual exclusion for the critical section.
func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money -= 10
		mutex.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}

	go stingy(&money, &mutex)
	go spendy(&money, &mutex)

	time.Sleep(2 * time.Second)

	// Protect read access with the mutex.
	// This ensures we see the most up-to-date value after all updates.
	mutex.Lock()
	fmt.Println("Money in bank account: ", money)
	mutex.Unlock()
}
