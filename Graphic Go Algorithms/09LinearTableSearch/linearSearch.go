package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	fmt.Printf("Please enter the value you want to search : \n")
	value := 0
	fmt.Scan(&value)

	isSearch := false
	length := len(scores)
	for i := 0; i < length; i++ {
		// กรณี value  มีใน scores 
		if scores[i] == value {
			isSearch = true
			fmt.Printf("Found value: %d the index is : %d", value, i)
			break
		}
	}
	// ไม่พบ value ใน scores
	if !isSearch {
		fmt.Printf("The value was not found : %d", value)
	}
}
