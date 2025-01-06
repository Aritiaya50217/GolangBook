package main

import "fmt"

func main() {
	// close
	valueStream := make(chan interface{})
	close(valueStream)

	intStream := make(chan int)
	fmt.Printf("Before : %v\n", intStream)

	close(intStream)
	integer, ok := <-intStream
	fmt.Println("----- Closed -----")
	fmt.Printf("(%v) : %v\n", ok, integer)
}
