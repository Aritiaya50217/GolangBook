package main

import (
	"fmt"
)

// Time Complexity: O(1) เนื่องจากการเข้าถึงค่าของอาร์เรย์ที่ดัชนีใดๆ จะใช้เวลาเท่ากันไม่ว่าขนาดของอาร์เรย์จะใหญ่แค่ไหน
func GetElement(arr []int, index int) int {
	return arr[index]
}

// Time Complexity: O(n) เพราะในกรณีที่แย่ที่สุด อัลกอริธึมต้องค้นหาผ่านทุกค่าในอาร์เรย์ ซึ่งจะใช้เวลาเพิ่มขึ้นตามจำนวนสมาชิกในอาร์เรย์
func LinearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i // คืนค่าดัชนีของค่าที่พบ
		}
	}
	return -1 // คืนค่า -1 ถ้าไม่พบค่า
}

// Time Complexity: O(n^2) เพราะในทุกๆ การวนลูปหลัก (i) จะมีการวนลูปย่อย (j) ที่จะเปรียบเทียบสมาชิกภายในอาร์เรย์ทั้งหมด
// ซึ่งทำให้การทำงานเพิ่มขึ้นตามกำลังสองของขนาดข้อมูล

// BubbleSort จัดเรียง slice ของตัวเลขในลำดับที่เพิ่มขึ้น
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

// Time Complexity: O(log n) เพราะในแต่ละขั้นตอนของการค้นหา ขนาดของข้อมูลจะถูกแบ่งครึ่ง (โดยการเลือกค่ากลาง)
// ซึ่งทำให้เวลาเพิ่มขึ้นในลักษณะลอการิธึม

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		fmt.Printf("low : %d , high : %d\n", low, high)
		mid := (low + high) / 2
		fmt.Println("mid : ", mid)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
			fmt.Printf("arr[%d] = %d < %d  \n", mid, low, target)
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	// constant time
	arr := []int{10, 20, 30, 40}
	fmt.Println(GetElement(arr, 2))

	// linear time
	arr2 := []int{1, 2, 3, 4, 5}
	target := 4
	index := LinearSearch(arr2, target)
	fmt.Println("Found at index : ", index)

	// quadratic time
	arr3 := []int{64, 34, 25, 12, 22, 11, 90}
	BubbleSort(arr3)
	fmt.Println("Sorted array:", arr)

	// logarithmic Time
	arr4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target2 := 7
	result := BinarySearch(arr4, target2)
	fmt.Println("Found at index:", result)
}
