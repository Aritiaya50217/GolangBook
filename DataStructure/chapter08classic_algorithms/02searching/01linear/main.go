package main

import "fmt"

func LinearSearch(elements []int, findElement int) bool {
	for _, element := range elements {
		if element == findElement {
			return true
		}
	}

	return false
}

func main() {
	elements := []int{1, 2, 3, 4, 5, 6}
	target := 6
	fmt.Println(LinearSearch(elements, target))
}
