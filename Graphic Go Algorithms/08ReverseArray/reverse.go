package main

import "fmt"

func reverse(arrays []int, length int) {
	middle := length / 2
	for i := 0; i <= middle; i++ {
		var temp = arrays[i]
		arrays[i] = arrays[length-i-1]
		arrays[length-i-1] = temp
	}
}

func main() {
	scores := []int{50, 60, 70, 80, 90}
	length := len(scores)

	reverse(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

/*
		i=0		 	 	middle=length/2			arrays[length-i-1]
	[	50			60			70			80			90	]

	swap 50 and 90 (arrays[i] = arrays[length-i-1])
	[	90			60			70			80			50	]

	swap 60 and 80
	[	90			80			70			60			50	]
	

*/