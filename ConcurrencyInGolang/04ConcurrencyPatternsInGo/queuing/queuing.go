package main

import (
	"fmt"
	"time"
)

func queueWithChannel() {
	queue := make(chan int, 5) // สร้าง channel

	// goroutine สำหรับเพิ่มข้อมูลใน queue
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Enqueued : ", i)
			queue <- i // เพิ่มข้อมูลใน queue
			time.Sleep(500 * time.Millisecond)
		}
		close(queue) // ปิด channel หลังส่งข้อมูล
	}()

	// อ่านข้อมูลจาก queue
	for item := range queue {
		fmt.Println("Dequeued : ", item)
		time.Sleep(1 * time.Second)
	}
}

// worker ประมวลผลจาก queue
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2 // ส่งผลลัพธ์กลับ
	}
}

func queueWithConcurrency() {
	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// สร้าง workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// ส่งงานเข้าสู่คิว
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// รับผลลัพธ์จากคิว
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result : ", <-results)
	}

	defer close(results)

}
func main() {
	queueWithConcurrency()
}
