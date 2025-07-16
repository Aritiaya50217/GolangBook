package main

import "fmt"

func main() {
	scores := []int{50, 65, 99, 87, 74, 63, 76, 100}
	length := len(scores)

	sort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func sort(array []int, length int) {
	// สร้าง slice ใหม่เพื่อเก็บค่าที่ sort
	temp := make([]int, length)
	mergeSort(array, temp, 0, length-1)

}

func mergeSort(array []int, temp []int, left int, right int) {
	if left < right {
		center := (left + right) / 2
		// left merge sort (ครึ่งซ้าย)
		mergeSort(array, temp, left, center)
		// right merge sort (ครึ่งขวา)
		mergeSort(array, temp, center+1, right)
		// merge two ordered arrays (รวม)
		merge(array, temp, left, center+1, right)
	}
}

/*
	Combine two ordered list into an ordered list
	temp : Temporary array
	left : Start the subscript on the left
	right : Start the subscript on the right
	rightEndIndex : End subscript on the right

*/
func merge(array []int, temp []int, left int, right int, rightEndIndex int) {
	// End subscript on the left
	leftEndIndex := right - 1
	// Starting from the left count
	tempIndex := left
	elementNumber := rightEndIndex - left + 1

	for {
		if left > leftEndIndex || right > rightEndIndex {
			break
		}
		if array[left] <= array[right] {
			temp[tempIndex] = array[left]
			tempIndex++
			left++
		} else {
			temp[tempIndex] = array[right]
			tempIndex++
			right++
		}
	}
	for {
		if left > leftEndIndex {
			break
		}
		// if there  is element on the left
		temp[tempIndex] = array[left]
		tempIndex++
		left++
	}

	for {
		if right > rightEndIndex {
			break
		}
		// if there  is element on the right
		temp[tempIndex] = array[right]
		tempIndex++
		right++
	}

	for i := 0; i < elementNumber; i++ {
		array[rightEndIndex] = temp[rightEndIndex]
		rightEndIndex--
	}
}

/*
	 left= 0		  center=(left+right)/2=3 (ตำแหน่งที่ 3 => 87)							right=7
		50			65			99			87			74			63			76			100


				split																	split
	50		65		99		87											74		63		76		100


		split				split												split				split
	50			65		99			87										74			63		76			100


   Merge sort (เรียง)	   Merge sort 											Merge sort 		 	Merge sort
	50			65		87			99										63			74		76			100

			Merge sort (เรียง)															Merge sort (เรียง)
	50			65		87			99										63			74		76			100

												Merge sort (เรียง)
				50			63			65			74			76			87			99			100

*/
