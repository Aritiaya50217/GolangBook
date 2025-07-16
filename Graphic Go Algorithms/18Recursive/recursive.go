package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		// recursively call yourself until the end of the return
		return factorial(n-1) * n
	}
}

func main() {
	n := 5
	fac := factorial(n)
	fmt.Printf("The factorial of 5 is : %d", fac)
}
