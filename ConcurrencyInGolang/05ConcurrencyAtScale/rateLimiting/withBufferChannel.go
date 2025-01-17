package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string) {
	for job := range jobs {
		fmt.Printf("Worker %d processing %s at %s\n", id, job, time.Now())
		time.Sleep(500 * time.Millisecond) // จำลองเวลาการทำงาน
	}
}

func main() {
	rateLimit := 2
	jobs := make(chan string, rateLimit)

	// เริ่ม worker
	for i := 1; i <= rateLimit; i++ {
		go worker(i, jobs)
	}

	requests := []string{"req1", "req2", "req3", "req4", "req5"}
	for _, req := range requests {
		jobs <- req                        // ส่งคำขอไปยัง worker
		time.Sleep(300 * time.Millisecond) // ควบคุมอัตราการส่ง
	}
	close(jobs)
}
