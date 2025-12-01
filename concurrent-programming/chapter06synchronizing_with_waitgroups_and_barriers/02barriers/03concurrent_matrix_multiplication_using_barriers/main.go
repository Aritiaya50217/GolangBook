package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Barrier struct {
	size      int // total number of participants in the barrier
	waitCount int // counter variable representing the number of currently suspened executions.
	cond      *sync.Cond
}

func NewBarrier(size int) *Barrier {
	condVar := sync.NewCond(&sync.Mutex{})     // creates new condition variable.
	return &Barrier{size: size, cond: condVar} // creates and returns reference to new barrier.
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount += 1 // increments the count variable by 1.

	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast() // if waitCount has reached the barrier size, resets waitCount and broadcasts on the condition variable
	} else {
		b.cond.Wait() // If waitCount hasn’t reached the barrier size, waits on the condition variable.
	}
	b.cond.L.Unlock()
}

const matrixSize = 3

func matrixMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ { // Iterates over every row
			sum := 0
			for i := 0; i < matrixSize; i++ { // Iterates over every cloumn
				sum += matrixA[row][i] * matrixB[i][col] // sums up each value of the row from A multiplied by each value of the column from B
			}
			result[row][col] = sum // updates the result matrix with the sum
		}
	}
}

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5 // For every row and column, assigns a random number between –5 and 4
		}
	}
}

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int, row int, barrier *Barrier) {
	for {
		barrier.Wait() // Waits on the barrier until the main() goroutine loads the matrices
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col] // calculates the result of the row in this goroutine
			}
			result[row][col] = sum // assigns the result to the correct row and column
		}
		barrier.Wait() // Waits on the barrier until every other row has been computed
	}
}

func main() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	barrier := NewBarrier(matrixSize + 1) // creates a new barrier with size of row goroutines + main() goroutine
	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, barrier) // creates a go routine per row , assigning the correct row numbers
	}

	for i := 0; i < 4; i++ {
		// loads up both matrices by randomly generating them
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)

		barrier.Wait() // releases the barrier so the goroutines can start their computions

		barrier.Wait() // waits until the goroutines finish their computations

		for i := 0; i < matrixSize; i++ {
			fmt.Println(matrixA[i], matrixB[i], result[i])
		}
		fmt.Println()
	}
}
