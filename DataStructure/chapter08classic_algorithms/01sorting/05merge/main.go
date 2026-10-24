package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createArray(num int) []int {
	array := make([]int, num)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < num; i++ {
		array[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return array
}

func MergeSorter(array []int) []int {
	if len(array) < 2 {
		return array
	}

	middle := (len(array)) / 2
	return JoinArrays(MergeSorter(array[:middle]), MergeSorter(array[middle:]))
}

func JoinArrays(leftArr []int, rightArr []int) []int {
	num, i, j := len(leftArr)+len(rightArr), 0, 0
	array := make([]int, num)

	for k := 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
	return array
}

func main() {
	elements := createArray(40)
	fmt.Println("\n Before Sorting \n\n", elements)
	fmt.Println("\n-After Sorting\n\n", MergeSorter(elements))
}
