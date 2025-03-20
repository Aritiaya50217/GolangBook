package main

import (
	"fmt"
)

// fibonacci method given k integer
func fibonacci(k int) int {
	if k <= 1 {
		return 1
	}
	return fibonacci(k-1) + fibonacci(k-2)
}

// merge function to combine two sorted slice
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	// merge the two sorted slices
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// append any remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// mergeSort function that divides the array recursively and merges the sorted results
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])  // Recursively sort the left half
	right := mergeSort(arr[mid:]) // Recursively sort right half

	return merge(left, right) // merge the sorted halves
}

func main() {
	m := 5
	for m = 0; m < 8; m++ {
		fib := fibonacci(m)
		fmt.Println(fib)
	}

	// merge
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array: ", arr)

	sortArr := mergeSort(arr)
	fmt.Println("Sorted array :", sortArr)
}
