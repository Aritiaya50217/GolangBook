package main

import "fmt"

func producer(ch chan int) {
	ch <- 42
}
func consumer(ch chan int) {
	x := <-ch
	fmt.Println(x)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
