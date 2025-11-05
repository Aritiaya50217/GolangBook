package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	value int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewSemaphore(initial int) *Semaphore {
	s := &Semaphore{value: initial}
	s.cond = sync.NewCond(&s.mu)
	return s
}

func (s *Semaphore) Acquire() {
	s.mu.Lock()
	for s.value == 0 {
		s.cond.Wait() // รอจนกว่าจะมีสิทธิ์ว่าง
	}
	s.value-- // ใช้สิทธิ์
	s.mu.Unlock()
}

func (s *Semaphore) Release() {
	s.mu.Lock()
	s.value++       // คืนสิทธิ์
	s.cond.Signal() // ปลุก goroutine ที่รออยู่
	s.mu.Unlock()
}

func main() {

	sem := NewSemaphore(3)
	for i := 1; i <= 6; i++ {
		go func(id int) {
			sem.Acquire()
			fmt.Println("Worker", id, "started")
			time.Sleep(1 * time.Second)
			fmt.Println("Worker", id, "finished")
			sem.Release()
		}(i)
	}
	time.Sleep(5 * time.Second)

}
