package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	fmt.Printf("Please enter the index to be deleted: \n")
	index := 0
	fmt.Scan(&index)

	length := len(scores)
	// create a new array
	tempArray := make([]int, length-1)
	for i := 0; i < length; i++ {
		// copy data in front of index to the front of array
		// ถ้า i มีค่าน้อยกว่า index ให้คงที่ตำแหน่งเดิม
		if i < index {
			tempArray[i] = scores[i]
		}
		// copy the array after index to the end of tempArray
		// ถ้า i มากกว่า index ให้เลื่อนมาด้านหน้า 1 ตำแหน่ง
		if i > index {
			tempArray[i-1] = scores[i]
		}
	}
	scores = tempArray
	for i := 0; i < length-1; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

/*
					   i = 2
	[	90		70		50		80		60		85	]

	After delete index = 2
	[	90		70		80		60		85	]

*/