package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int)
	
	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	
	close(jobs)

	time.Sleep(6 * time.Second)

}
