package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Printf("Please enter the number of month : \n")
	var number int
	fmt.Scanf("%d", &number)

	for i := 1; i <= number; i++ {
		fmt.Printf("%d month: %d \n", i, fibonacci(i))
	}
}

/* Fibonacci definition
if n = 0 , 1
	fn = n
if n > 1
	fn = fn-1 + fn-2

1 month						 [x]						1
							  |
2 month						 [x]						1
						  /		 \
3 month					[x]		  [x]					2
					 /		\		  \
4 month 		   [x]		[x]		    [x]				3
				 /	  \		   \	    /  \
5 month		  [x]	  [x]	   [x]     [x]	[x]			5
			  / \	   |	  /   \	   /  \		\
6 month		[x]	[x]   [x]	 [x]  [x] [x] [x]	[x]		8
*/
