package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	length := len(scores)
	tempArray := make([]int, length+1)
	// insert 75 into the index = 2
	insert(scores, length, tempArray, 75, 2)

	scores = tempArray
	for i := 0; i < length+1; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func insert(array []int, length int, tempArray []int, scores int, insertIndex int) {
	for i := 0; i < length; i++ {
		if i < insertIndex {
			tempArray[i] = array[i]
		} else {
			tempArray[i+1] = array[i]
		}
	}
	tempArray[insertIndex] = scores
}

/*
	[	90		70		50		80		60		85	]
					insert i = 2
	[	90		70		75		50		80		60		85	]

*/
