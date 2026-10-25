package main

import "fmt"

func Interpolationsearch(elements []int, element int) (bool, int) {
	high := len(elements) - 1
	low := 0
	mid := 0

	for elements[low] < element && elements[high] > element {
		mid = low + ((element-elements[low])*(high-low))/(elements[high]-elements[low])
		if elements[mid] < element {
			low = mid + 1
		} else if elements[mid] > element {
			high = mid - 1
		} else {
			return true, mid
		}
	}

	if elements[low] == element {
		return true, low
	} else if elements[high] == element {
		return true, high
	}

	return false, -1
}

func main() {
	elements := []int{2, 3, 5, 7, 9}
	element := 7

	found, index := Interpolationsearch(elements, element)
	fmt.Printf("%v found at %d\n", found, index)
}
