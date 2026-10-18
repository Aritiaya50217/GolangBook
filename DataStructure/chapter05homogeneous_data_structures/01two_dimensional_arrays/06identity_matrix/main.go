package main

import "fmt"

func Identity(order int) [][]float64 {
	matrix := make([][]float64, order)
	for i := 0; i < order; i++ {
		temp := make([]float64, order)
		temp[i] = 1
		matrix[i] = temp
	}
	return matrix
}

func main() {
	fmt.Println("Identity : ", Identity(4))
}
