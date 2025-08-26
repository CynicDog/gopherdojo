package main

import (
	"fmt"

	"github.com/cynicdog/gopherdojo/concurrent-programming/ch05/semaphore"
)

func doWork(semaphore *Semaphore) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	semaphore.Release()
}

func main() {
	semaphore := NewSemaphore(0)
	for i := 0; i < 50; i++ {
		go doWork(semaphore)
		fmt.Println("Waiting for child goroutine")
		semaphore.Acquire()
		fmt.Println("Child goroutine finished")
	}
}
