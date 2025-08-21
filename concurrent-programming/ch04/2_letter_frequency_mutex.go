package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)

	// Lock the mutex once before the loop instead of inside the loop.
	// This critical section ensures that only one goroutine modifies the shared frequency slice at a time.
	// Placing the lock outside the loop avoids repeatedly locking and unlocking for every character,
	// which would be much less efficient due to mutex overhead.
	// By keeping the lock for the whole batch update, we reduce contention while still protecting shared data.
	mutex.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	mutex.Unlock()
	fmt.Println("Completed: ", url, time.Now().Format("15:04:05"))
}

func main() {
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	time.Sleep(5 * time.Second)

	// Lock mutex before reading shared data to ensure memory visibility
	mutex.Lock()
	for i, c := range allLetters {
		fmt.Printf("%c-%d\n", c, frequency[i])
	}
	mutex.Unlock()
}
