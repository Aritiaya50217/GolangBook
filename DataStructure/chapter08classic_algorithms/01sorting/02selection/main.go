package main

import "fmt"

func selectionSort(numbers []int) {
	n := len(numbers)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}
}

func main() {
	numbers := []int{5, 3, 1, 8, 9, 10}
	selectionSort(numbers)
	fmt.Println("selectionSort : ", numbers)
}
