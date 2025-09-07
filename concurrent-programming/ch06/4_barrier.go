package main

import (
	"fmt"
	"time"

	"github.com/cynicdog/gopherdojo/concurrent-programming/ch06/barrier"
)

func workAndWait(name string, timeToWork int, barrier *barrier.Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait()
	}
}

func main() {
	barrier := barrier.New(2) 
	go workAndWait("Red", 4, barrier)
	go workAndWait("Blue", 10, barrier)
	time.Sleep(100 * time.Second)
}