package main

import (
	"fmt"
)

func Slice() {
	var mySlice []int
	mySlice = append(mySlice, 1, 2, 3, 4)
	firstElement := mySlice[0]
	fmt.Println("firstElement : ", firstElement)

	subset := mySlice[1:4]
	fmt.Println("subset : ", subset)

	subset = mySlice[1:]
	fmt.Println("mySlice[1:] => ", subset)

	subset = mySlice[0 : len(mySlice)-1]
	fmt.Println(subset)

	// To copy a slice, use the copy() function.
	// If you assign one slice to another with the equal operator,
	// the slices will point at the same memory location,
	// and changing one would change both slices.
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 4)

	// Create a unique copy in memory
	copy(slice2, slice1)

	// Changing one should not affect the other
	slice2[3] = 99
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func main() {
	mySlice := make([]byte, 8, 128)
	fmt.Println("max : ", cap(mySlice))
	// Current length of slice
	fmt.Println("Length:", len(mySlice))

	Slice()
}
