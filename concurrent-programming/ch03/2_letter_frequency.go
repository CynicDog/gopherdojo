package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

// countLetters counts letter frequencies from a URL.
// The 'frequency' slice header is passed by value (copied),
// but it points to the same underlying array, so updates are visible
// to the caller (or in this case, shared across goroutines!).
func countLetters(url string, frequency []int) {

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			// Update the underlying array.
			// This is shared among all goroutines, so concurrent writes
			// may cause **race conditions**.
			frequency[cIndex] += 1
		}
	}

	fmt.Println("Completed", url)
}

func main() {

	// Slice header (frequency) lives on the stack of main().
	// The underlying array may escape to the heap because it is shared with goroutines.
	var frequency = make([]int, 26)

	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		// Start countLetters in a goroutine
		// Each goroutine gets a copy of the slice header pointing to the same underlying array
		go countLetters(url, frequency)
	}
	time.Sleep(5 * time.Second)
	for i, c := range allLetters {
		fmt.Printf("%c-%d\n", c, frequency[i])
	}
}
