package main

import (
	"fmt"
	"sync"
	"time"
)

func _receiver(messages chan int, wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = <-messages
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Received:", msg)
	}
	wGroup.Done()
}

func main() {

	// Creates a buffered channel and feeds messages into the channel
	// at faster rate than the receiver can consume them
	msgChannel := make(chan int, 3)

	wGroup := sync.WaitGroup{}
	wGroup.Add(1)

	go _receiver(msgChannel, &wGroup)

	for i := 1; i <= 6; i++ {
		size := len(msgChannel)
		fmt.Printf(
			"%s Sending: %d. Buffer Size: %d\n",
			time.Now().Format("2006-01-02 15:04:05"),
			i,
			size,
		)
		msgChannel <- i
	}
	msgChannel <- -1
	wGroup.Wait()
}
