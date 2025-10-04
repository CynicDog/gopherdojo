package main

import "fmt"

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	someBigNumber := 3419110
	resultChannelFirst := make(chan []int)

	anotherBigNumber := 4033836
	resultChannelSecond := make(chan []int)

	// Launch an anonymous goroutine
	go func() {
		// Compute factors and send the result into the channel
		resultChannelFirst <- findFactors(someBigNumber)
	}()

	go func() {
		resultChannelSecond <- findFactors(anotherBigNumber)
	}()

	// Receive and print the result from each goroutine
	fmt.Println(<-resultChannelFirst)
	fmt.Println(<-resultChannelSecond)
}
