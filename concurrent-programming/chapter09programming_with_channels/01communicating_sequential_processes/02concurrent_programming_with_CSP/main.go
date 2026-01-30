package main

import "fmt"

func producer(ch chan int) {
	ch <- 100
}

func consumer(ch chan int) {
	x := <-ch
	fmt.Println(x)
}

func generate(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- 1
	}
	close(ch)
}

func square(in, out chan int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)

	fmt.Println("-- Pipeline (CSP Pattern) --")
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generate(ch1)
	go square(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}
}
