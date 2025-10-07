// channels_pipeline.go
//
// This example demonstrates Go's idiomatic concurrency pattern:
// "Don't communicate by sharing memory; share memory by communicating."
//
// It shows a simple pipeline of goroutines connected by channels.
// Each goroutine receives data from its input channel, processes it,
// and sends the result to the next stage via an output channel.
//
// This approach avoids explicit locking (mutexes) by using channels
// to synchronize access and communicate values safely.

package main

import (
	"fmt"
	"strings"
)

// stage1 generates a stream of messages and sends them to the next stage.
func stage1(out chan<- string) {
	defer close(out)
	for _, msg := range []string{"hello", "concurrency", "world"} {
		out <- msg
	}
}

// stage2 receives strings, converts them to uppercase, and forwards them.
func stage2(in <-chan string, out chan<- string) {
	defer close(out)
	for msg := range in {
		out <- strings.ToUpper(msg)
	}
}

// stage3 receives messages and prints them (final stage).
func stage3(in <-chan string) {
	for msg := range in {
		fmt.Println("Received:", msg)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Launch each stage as a separate goroutine.
	go stage1(ch1)      // Produces messages
	go stage2(ch1, ch2) // Transforms messages
	stage3(ch2)         // Consumes and prints messages

	// Note: Sending and receiving pauses each goroutine until
	// the other side is ready — providing safe synchronization.
	//
	// No mutexes, no shared global variables — just clear data flow.
	//
	// Output (order guaranteed because of sequential channel flow):
	//   Received: HELLO
	//   Received: CONCURRENCY
	//   Received: WORLD
}
