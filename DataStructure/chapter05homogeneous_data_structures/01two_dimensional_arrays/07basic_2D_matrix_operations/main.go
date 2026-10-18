package main

import (
	"fmt"
)

// add method
func add(matrix1, matrix2 [2][2]int) [2][2]int {
	var sum [2][2]int
	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

// subtract method
func subtract(matrix1, matrix2 [2][2]int) [2][2]int {
	var difference [2][2]int
	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference
}

// multiply method
func multiply(matrix1, matrix2 [2][2]int) [2][2]int {
	var product [2][2]int
	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			var productSum int = 0
			for n := 0; n < 2; n++ {
				productSum = productSum + matrix1[l][n]*matrix2[n][m]
				fmt.Println("matrix1[l][n] => ", matrix1[l][n])
				fmt.Println("matrix2[n][m] => ", matrix2[n][m])
				fmt.Println("value : ", matrix1[l][n]*matrix2[n][m])
				fmt.Println("productSum : ", productSum)
				fmt.Println("productSum + value : ", productSum+matrix1[l][n]*matrix2[n][m])
			}
			product[l][m] = productSum
		}
	}
	return product
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var transMatrix [2][2]int
	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

func determinant(matrix1 [2][2]int) float64 {
	det := float64(matrix1[0][0]*matrix1[1][1]) - float64(matrix1[0][1]*matrix1[1][0])
	// fmt.Println("float64(matrix1[0][0]*matrix1[1][1] : ", float64(matrix1[0][0]*matrix1[1][1]))
	// fmt.Println("float64(matrix1[0][1]*matrix1[1][0]) : ", float64(matrix1[0][1]*matrix1[1][0]))
	// fmt.Println("float64(matrix1[0][0]*matrix1[1][1]) - float64(matrix1[0][1]*matrix1[1][0]) : ", det)
	return det
}

// inverse method
func inverse(matrix [2][2]int) [2][2]float64 {
	det := determinant(matrix)
	var invmatrix [2][2]float64
	invmatrix[0][0] = float64(matrix[1][1]) / det
	invmatrix[0][1] = -float64(matrix[0][1]) / det
	invmatrix[1][0] = -float64(matrix[1][0]) / det
	invmatrix[1][1] = float64(matrix[0][0]) / det
	return invmatrix
}

func main() {
	var matrix1 = [2][2]int{
		{4, 5},
		{1, 2},
	}
	var matrix2 = [2][2]int{
		{6, 7},
		{3, 4},
	}

	fmt.Println("add : ", add(matrix1, matrix2))
	fmt.Println("subtract : ", subtract(matrix1, matrix2))
	fmt.Printf("\n===== multiply =====\n")
	fmt.Println("\n\nmultiply : ", multiply(matrix1, matrix2))
	fmt.Println("transpose : ", transpose(matrix1))
	fmt.Printf("\n===== determinant =====\n")
	fmt.Println("determinant : ", determinant(matrix1))
	fmt.Printf("\n===== inverse =====\n")
	fmt.Println("inverse : ", inverse(matrix1))
}
