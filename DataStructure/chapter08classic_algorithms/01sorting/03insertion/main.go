package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSequence(num int) []int {
	sequence := make([]int, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		sequence[i] = rand.Intn(999) - rand.Intn(999)
	}
	return sequence
}

func InsertionSorter(elements []int) {
	for i := 1; i < len(elements); i++ {
		j := i
		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j - 1
		}
	}
}

func main() {
	sequence := randomSequence(24)
	fmt.Println("\n====== Before Sorting === \n\n", sequence)
	InsertionSorter(sequence)
	fmt.Println("\n====== After Sorting ======\n\n", sequence)
}
