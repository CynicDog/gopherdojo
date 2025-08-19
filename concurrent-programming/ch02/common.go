package main

import (
	"fmt"
	"time"
)

func DoWork(id int) {
	fmt.Printf("Work %d started at %s\n", id, time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("Work %d finished at %s\n", id, time.Now())
}
