package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{1, 3, 16, 10, 45, 31, 28, 36, 45, 75}
	element := 36

	i := sort.Search(len(elements), func(i int) bool {
		return elements[i] >= element
	})

	if i < len(elements) && elements[i] == element {
		fmt.Printf("found element %d at index %d in %v\n", element, i, elements)
	} else {
		fmt.Printf("element %d not found in %v\n", element, elements)
	}
}
