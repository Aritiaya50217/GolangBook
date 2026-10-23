package main

import "fmt"

func bubbleSorter(numbers []int) {
	num := len(numbers)
	for i := 0; i < num; i++ {
		for j := 0; j < num-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func main() {
	numbers := []int{5, 3, 1, 8, 9, 1}
	bubbleSorter(numbers)
	fmt.Println("bubbleSorter : ", numbers)
}
