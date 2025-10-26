package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d running on OS thread %d\n", id, runtime.LockOSThread)
}

func main() {
	// Go defaults the value of GOMAXPROCS to the value of NumCPU()
	fmt.Println("Number of CPUs : ", runtime.NumCPU())
	// Calling GOMAXPROCS(n) with n < 1 returns the current value
	fmt.Println("GOMAXPROCS : ", runtime.GOMAXPROCS(0))

	fmt.Println("\n== worker ==")
	runtime.GOMAXPROCS(4) // ใช้ได้สูงสุด 4 CPU core
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		/*
			Go runtime จะกระจาย goroutines 10 ตัวให้รันบน 4 threads
			OS เห็นแค่ 4 threads, แต่ Go runtime สลับ goroutines ภายใน user space เอง 
		*/
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines completed")
}
