package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string) {
	for job := range jobs {
		fmt.Printf("Worker %d processing %s\n", id, job)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan string, 10)

	// start workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	// push jobs
	for i := 1; i <= 5; i++ {
		jobs <- fmt.Sprintf("Job-%d", i)
	}
	close(jobs)

	time.Sleep(5 * time.Second)

}
