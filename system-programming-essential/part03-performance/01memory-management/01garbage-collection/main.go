package main

import (
	"fmt"
)

func createData() {
	data := make([]int, 1000000)
	fmt.Println("data created : ", len(data))
}

func main() {
	createData()
	fmt.Println("function finished")
}
