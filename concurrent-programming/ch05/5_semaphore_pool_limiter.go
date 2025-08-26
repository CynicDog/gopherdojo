package main

import (
	"fmt"

	"github.com/cynicdog/gopherdojo/concurrent-programming/ch05/semaphore"
)

func _doWork(semaphore *Semaphore, id int, wg *sync.WaitGroup) {
	defer wg.Done()           // mark this worker as finished when function exits
	semaphore.Acquire()       // request a permit (may block if none available)
	defer semaphore.Release() // return the permit when finished

	fmt.Printf("[Worker %d] started\n", id)
	fmt.Printf("[Worker %d] finished\n", id)
}

func main() {
	const maxConcurrent = 3
	semaphore := NewSemaphore(maxConcurrent)

	// wg (WaitGroup) is used to wait until all launched goroutines finish
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go _doWork(semaphore, i, &wg)
	}
	wg.Wait() // wait for all workers to finish
}
