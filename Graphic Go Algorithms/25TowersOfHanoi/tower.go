package main

import "fmt"

func hanoi(n int, A string, B string, C string) {
	if n == 1 {
		fmt.Printf("Move %d %s to %s \n", n, A, C)
	} else {
		// Move the n-1th disc on the A through C to B
		hanoi(n-1, A, C, B)
		fmt.Printf("Move %d from %s to %s \n", n, A, C)
		// Move the n-1th disc on the B through A to C
		hanoi(n-1, B, A, C)
	}
}

func main() {
	fmt.Printf("Please enter number of discs : \n")
	var n int
	fmt.Scanf("%d", &n)
	hanoi(n, "A", "B", "C")
}

/*
Towers of Hanoi
1. If there is only one disc , move it direcly to C (A -> C)
2. When there are two dissc , use B as an auxiliary(A -> B , A -> C , B -> C)
3. If there are more than two discs, use B as an auxiliary(A -> B, A -> C , B-> C) ,
and coutinue to recursive process.

*/
