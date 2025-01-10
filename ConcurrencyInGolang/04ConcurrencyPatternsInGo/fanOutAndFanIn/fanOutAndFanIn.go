package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(time.Millisecond * 100)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Fan-Out สร้าง workers หลายตัว
	for w := 1; w <= 3; w++ {
		go workers(w, jobs, results)
	}

	// ส่งงานไปยัง Channel
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Fan-In รวบรวมผลลัพธ์
	for a := 1; a <= 9; a++ {
		fmt.Println("Result:", <-results)
	}
}
