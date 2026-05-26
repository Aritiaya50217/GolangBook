package main

import "fmt"

func worker(ch chan int) {
	num := <-ch
	fmt.Println("Worker got : ", num)
}

func main() {
	ch := make(chan int)
	go worker(ch)
	ch <- 42
}
