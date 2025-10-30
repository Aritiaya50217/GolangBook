package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID int
}

func worker(id int, jobs <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing Task %d\n", id, job.ID)
	}
}

func main() {
	jobs := make(chan Task, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- Task{ID: j}
	}
	close(jobs)
	wg.Wait()
}
