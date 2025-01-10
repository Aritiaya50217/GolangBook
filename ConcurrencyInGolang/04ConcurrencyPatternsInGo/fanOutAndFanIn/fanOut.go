package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // จำลองการประมวลผล
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func fanOut() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// สร้าง  Workers (Fan-Out)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// ส่งงานไปยัง Channel
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)

	// รับ result จาก workers
	for i := 1; i <= 3; i++ {
		fmt.Println("Result:", <-results)
	}

}
func main() {
	fanOut()
}
