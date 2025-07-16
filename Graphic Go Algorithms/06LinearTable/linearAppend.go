package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	fmt.Printf("Before : %d\n", scores)
	length := len(scores)
	// create a new array
	// length+1 => ลำดับถัดไป
	tempArray := make([]int, length+1)
	for i := 0; i < length; i++ {
		tempArray[i] = scores[i]
	}
	// assign 75 
	tempArray[length] = 75

	scores = tempArray
	for i := 0; i < length+1; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

/*
	[	90	70	50	80	60	85	]
	copy
								assign 75
	[	90	70	50	80	60	85	75]

*/
