package main

import "fmt"

func increase(n *int) {
	*n++
}

func main() {
	x := 10

	increase(&x)

	fmt.Println(x)
}
