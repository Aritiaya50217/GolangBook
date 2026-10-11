package main

import "fmt"

func main() {
	rows := 7
	cols := 9
	twoSlices := make([][]int, rows)
	for i := range twoSlices {
		twoSlices[i] = make([]int, cols)
	}
	fmt.Println(twoSlices)

	test := make([]int, 0)
	test = append(test, 1, 2)
	fmt.Println(len(test), cap(test))
}
