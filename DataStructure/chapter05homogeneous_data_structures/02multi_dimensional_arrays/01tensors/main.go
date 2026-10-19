package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var array [3][3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				array[i][j][k] = rand.Intn(3)
				fmt.Printf("%d ", array[0][j][k]) // zero mode unflold
				fmt.Printf("%d ", array[1][j][k]) // one mode unflold
				fmt.Printf("%d ", array[2][j][k]) // two mode unflold
			}
			fmt.Printf("\n")
		}
	}
}
