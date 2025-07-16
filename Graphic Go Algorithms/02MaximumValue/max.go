package main

import "fmt"

func max(array []int, length int) int {
	// length-1 => ลำดับถัดจาก i
	for i := 0; i < length-1; i++ {
		// array[i+1] => เลขก่อนหน้า
		// array[i] => เลขที่ใช้เทียบ
		if array[i] > array[i+1] {
			// swap => สลับตำแหน่ง เลขที่มากกว่าจะไปด้านหลัง
			var temp = array[i]
			array[i] = array[i+1]
			array[i+1] = temp
		}
	}
	// เลขมากสุดที่อยู่ลำดับท้าย
	var maxValue = array[length-1]
	return maxValue
}

func main() {
	var scores = []int{60, 70, 95, 50, 80}
	length := len(scores)
	maxValue := max(scores, length)
	fmt.Printf("Max Value : %d\n", maxValue)

}
