package main

import "fmt"

func Factor(num int) int {
	if num <= 1 {
		return 1
	}
	return num * Factor(num-1)
}

func main() {
	num := 12
	fmt.Printf("Factorial: %d is %d\n", num, Factor(num))
}
