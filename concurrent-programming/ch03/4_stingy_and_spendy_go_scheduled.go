package main

import (
	"fmt"
	"runtime"
	"time"
)

func scheduledStingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
		runtime.Gosched()
	}
	fmt.Println("Stingy Done")
}

func scheduledSpendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
		runtime.Gosched()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	go scheduledStingy(&money)
	go scheduledSpendy(&money)

	time.Sleep(2 * time.Second)
	println("Money in the bank account: ", money)
}
