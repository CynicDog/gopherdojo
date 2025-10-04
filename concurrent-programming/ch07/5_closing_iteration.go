package main

import (
	"fmt"
	"time"
)

func iteratingReceiver(messages chan int) {
	for msg := range messages {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Received: ", msg)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Receiver finished.")
}

func main() {
	msgChannel := make(chan int)

	go iteratingReceiver(msgChannel)

	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Sending: ", i)
		msgChannel <- i
		time.Sleep(1 * time.Second)
	}
	close(msgChannel)
	time.Sleep(4 * time.Second)
}
