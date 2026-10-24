package main

import "fmt"

func ShellSorter(elements []int) {
	var (
		n         = len(elements)
		intervals = []int{1}
		k         = 1
	)

	for {
		interval := power(2, k) + 1
		if interval > n-1 {
			break
		}
		intervals = append(intervals, intervals...)
		k++
	}

	for _, interval := range intervals {
		for i := interval; i < n; i += interval {
			j := i
			for j > 0 {
				if elements[j-interval] > elements[j] {
					elements[j-interval], elements[j] = elements[j], elements[j-interval]
				}
				j = j - interval
			}
		}
	}
}

func power(exponent int, index int) int {
	power := 1
	for index > 0 {
		if index&1 != 0 {
			power *= exponent
		}
		index >>= 1
		exponent *= exponent
	}
	return power
}

func main() {
	elements := []int{34, 202, 13, 19, 6, 5, 1, 43, 506, 12, 20, 28, 17, 100, 25, 4, 5, 97, 1000, 27}
	ShellSorter(elements)
	fmt.Println("elements : ", elements)
}
