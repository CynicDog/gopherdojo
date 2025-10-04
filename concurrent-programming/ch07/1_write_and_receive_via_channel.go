package main

import "fmt"

func main() {

	// Go's channels are synchronous by default
	msgChannel := make(chan string)

	go receiver(msgChannel)

	fmt.Println("Sending HELLO ..")
	msgChannel <- "HELLO"

	fmt.Println("Sending WORLD ..")
	msgChannel <- "WORLD"

	fmt.Println("Sending STOP ..")
	msgChannel <- "STOP"
}

func receiver(messages chan string) {
	msg := ""
	for msg != "STOP" {
		msg = <-messages
		fmt.Println("Received: ", msg)
	}
}
