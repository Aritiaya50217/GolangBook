package main

import "fmt"

func min(arrays []int, length int) int {
	// the index of the minimum
	minIndex := 0
	for j := 1; j < length; j++ {
		// ถ้าตัวเลขที่ใช้เทียบมากกว่า ให้เปลี่ยนตัวเทียบเป็นตัวที่น้อยกว่า แล้วเทียบไปเรื่อยๆ
		if arrays[minIndex] > arrays[j] {
			minIndex = j
		}
	}
	return arrays[minIndex]

	/* 
		minIndex=0  j=1 (ถ้า 60 < 80 ให้เทียบตัวถัดไป)
		[ 	60		80		95		50		70	]
		minIndex=0  		j=2 (ถ้า 60 < 95 ให้เทียบตัวถัดไป)
		[ 	60		80		95		50		70	]
		minIndex=0  				j=3 (ถ้า 60 > 50 ให้เทียบตัวถัดไป กรณีนี้ 50 น้อยกว่าจึงเปลี่ยนตัวเทียบ)
		[ 	60		80		95		50		70	]
								minIndex=0 j=4 (ถ้า 50 < 70 ให้เทียบตัวถัดไป)
		[ 	60		80		95		50		70	]

	Min Value : 50
	*/
}
func main() {
	scores := []int{60, 80, 95, 50, 70}
	length := len(scores)
	minValue := min(scores, length)
	fmt.Printf("Min Value : %d\n", minValue)
}
