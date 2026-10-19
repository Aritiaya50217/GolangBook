package main

import "fmt"

func PrintSpiral(n int) []int {
	left := 0
	top := 0
	right := n - 1
	bottom := n - 1
	size := n * n
	s := make([]int, size)
	i := 0

	for left < right {
		for c := left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++

		for r := top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--

		if top == bottom {
			break
		}

		for c := right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--

		for r := bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	s[top*n+left] = i

	return s
}

func main() {
	n := 5
	length := 2
	for i, sketch := range PrintSpiral(n) {
		fmt.Printf("%*d ", length, sketch)
		if i%n == n-1 {
			fmt.Println("")
		}
	}
}
