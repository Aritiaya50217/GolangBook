package main

import "fmt"

// twiceValue method given slice of int type
func twiceValue(slice []int) {
	for i, value := range slice {
		slice[i] = 2 * value
	}
}

func main() {
	slice := []int{1, 3, 5, 6}
	twiceValue(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println("new slice value", slice[i])
	}
}
