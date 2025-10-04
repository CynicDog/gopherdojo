package main

import (
	"fmt"
	"time"
)

func __receiver(messages chan int) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Received ", msg)
	}
}

func sender(message chan<- int) {
	for i := 1; ; i++ {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Sending ", i)
		message <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	msgChannel := make(chan int)
	go __receiver(msgChannel)
	go sender(msgChannel)

	time.Sleep(5 * time.Second)
}
