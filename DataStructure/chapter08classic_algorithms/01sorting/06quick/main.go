package main

import "fmt"

func QuickSorter(elements []int, below, upper int) {
	if below < upper {
		part := divideParts(elements, below, upper)
		QuickSorter(elements, below, part-1)
		QuickSorter(elements, part+1, upper)
	}
}

func divideParts(elements []int, below int, upper int) int {
	center := elements[upper]
	i := below

	for j := below; j < upper; j++ {
		if elements[j] <= center {
			swap(&elements[i], &elements[j])
			i += 1
		}
	}
	swap(&elements[i], &elements[upper])
	return i
}

func swap(element1, element2 *int) {
	val := *element1
	*element1 = *element2
	*element2 = val
}

func main() {
	var num int
	fmt.Print("Enter Number of Elements: ")
	fmt.Scan(&num)

	array := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Print("array[", i, "] : ")
		fmt.Scan(&array[i])
	}
	
	fmt.Print("Elements: ", array, "\n")
	QuickSorter(array, 0, num-1)
	fmt.Print("Sorted Elements: ", array, "\n")
}
