package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}
func main() {
	squere, cube := powerSeries(3)
	fmt.Println("Sequere : ", squere, "Cube : ", cube)

}
