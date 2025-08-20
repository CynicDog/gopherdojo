package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("Hello")
}

// Step 0: Program entry point.
// 1. The Go tool compiles this file (and any imports) into an executable.
// 2. The Go runtime is linked into the program automatically.
//   - The runtime sets up things like memory management (GC),
//     goroutine scheduling, and communication with the OS.
//
// 3. The runtime creates a small pool of kernel threads (M’s).
// 4. The runtime starts scheduling goroutines (G’s) onto those threads.
// 5. Finally, it calls func main() — which is the true entry point of our program.
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
