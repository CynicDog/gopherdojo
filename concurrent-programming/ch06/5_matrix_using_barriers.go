package main

import (
	"fmt"
	"math/rand"

	"github.com/cynicdog/gopherdojo/concurrent-programming/ch06/barrier"
)

const matrixSize = 3

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func rowMultiply(
	matrixA,
	matrixB,
	result *[matrixSize][matrixSize]int,
	row int,
	barrier *barrier.Barrier,
) {
	for {
		// Worker waits here until the main goroutine signals
		// that new matrices A and B have been generated and are ready.
		barrier.Wait()

		// Perform multiplication for one row of the result matrix.
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}

		// Worker signals it has finished computing this row.
		// Then waits until all other workers (and the main goroutine)
		// also finish, so the result matrix is complete and consistent.
		barrier.Wait()
	}
}

func main() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	barrier := barrier.New(matrixSize + 1)

	// Launch one worker per row of the result matrix
	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, barrier)
	}

	for i := 0; i < matrixSize+1; i++ {
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)

		// Main goroutine joins the workers at the barrier,
		// releasing them all to start computing their rows
		// once matrices A and B are ready.
		barrier.Wait()

		// Main goroutine waits here until all workers finish
		// computing their rows. At this point, the result
		// matrix is fully populated and safe to print.
		barrier.Wait()

		for i := 0; i < matrixSize; i++ {
			fmt.Println(matrixA[i], matrixB[i], result[i])
		}
		fmt.Println()
	}
}
