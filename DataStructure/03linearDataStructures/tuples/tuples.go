package main

import "fmt"

// h finction which returns the product of parameters x and y
func h(x, y int) int {
	return x * y
}

// g function which returns a and y parameters after  modification
func g(l, m int) (x, y int) {
	x = 2 * l
	y = 4 * m
	return x, y
}
func main() {

	fmt.Println(h(2, 2))
	fmt.Println(h(g(2,4)))
}
