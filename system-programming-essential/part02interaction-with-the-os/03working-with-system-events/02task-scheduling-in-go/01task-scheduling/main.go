package main

import (
	"fmt"
	"time"
)

type Job func()

type Scheduler struct {
	jobQueue chan Job
}

// NewScheduler creates a new Scheduler
func NewScheduler(size int) *Scheduler {
	return &Scheduler{jobQueue: make(chan Job, size)}
}

func (s *Scheduler) Start() {
	for job := range s.jobQueue {
		go job() // run the job in a new goroutine
	}
}

// schedule a job to be executed after a delay
func (s *Scheduler) Scheduler(job Job, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		s.jobQueue <- job
	}()
}

func main() {
	scheduler := NewScheduler(10) // buffer size of 10

	// schedule a job to run after 5 seconds
	scheduler.Scheduler(func() {
		fmt.Println("Job executed at ", time.Now())
	}, 5*time.Second)

	// start the scheduler
	go scheduler.Start()

	// wait for input to exit
	fmt.Println("Scheduler started. Press Enter to exit.")
	fmt.Scanln()
}
