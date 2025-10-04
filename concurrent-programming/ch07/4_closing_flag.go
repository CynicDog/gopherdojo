package main

import (
	"fmt"
	"time"
)

func closingReceiver(messages chan int) {
	for {
		msg, more := <-messages

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), msg, more)
		time.Sleep(1 * time.Second)
		if !more {
			return
		}
	}
}

func main() {
	msgChannel := make(chan int)

	go closingReceiver(msgChannel)

	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Sending: ", i)
		msgChannel <- i
		time.Sleep(1 * time.Second)
	}
	close(msgChannel)
	time.Sleep(5 * time.Second)
}
