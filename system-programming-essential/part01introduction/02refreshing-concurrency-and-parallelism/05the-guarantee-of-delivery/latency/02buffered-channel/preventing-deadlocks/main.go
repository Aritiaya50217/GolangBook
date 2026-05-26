package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	ch <- "Hello"

	go func() {
		fmt.Println("received : ", <-ch)
	}()

	fmt.Scanln()
}
