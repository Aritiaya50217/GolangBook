package main

import (
	"fmt"
	"time"
)

func deadlock(ch chan int) {
	ch <- 1 // จะค้างถ้าไม่มี goroutine อื่นรับ
	fmt.Println("This will not be printed")
}

func main() {
	ch := make(chan int)
	go deadlock(ch)

	select {
	case val := <-ch:
		fmt.Println("Received value : ", val)
	case <-time.After(1 * time.Second): // Timeout เพื่อป้องกัน Deadlock
		fmt.Println("Deadlock detected")

	}
}
