package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	// Step 1: main goroutine launches a new goroutine.
	// This registers `sayHello` with the Go scheduler,
	// so the scheduler now knows there’s another goroutine ready to run.
	go sayHello()

	// Step 2: main calls runtime.Gosched().
	// This is main explicitly yielding to the scheduler, saying:
	// "Pause me and let another goroutine have the CPU."
	runtime.Gosched()

	// Step 3: the scheduler picks `sayHello` (since it's runnable) and runs it,
	// which prints "Hello". When `sayHello` finishes,
	// the scheduler resumes the main goroutine.

	// Step 4: main continues and prints "Finished".
	fmt.Println("Finished")
}
