package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- 1
		fmt.Println("Produced ", i)
	}
}

func consumer(ch chan int) {
	for item := range ch {
		time.Sleep(time.Second)
		fmt.Println("Consumed ", item)
	}
}

func main() {
	ch := make(chan int, 3)
	fmt.Println("cap of ch : ", cap(ch))
	
	go producer(ch)
	go consumer(ch)

	time.Sleep(7 * time.Second)
}
