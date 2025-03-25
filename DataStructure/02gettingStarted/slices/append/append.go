package main

import "fmt"

func main() {
	arr := []int{5, 6, 7, 8, 9}
	slice1 := arr[:3]
	fmt.Println("slice1 : ", slice1)
	slice2 := arr[1:5]
	fmt.Println("slice2 : ", slice2)
	slice3 := append(slice2, 12)
	fmt.Println("slice2 : ", slice3)

}
