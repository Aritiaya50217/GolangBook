package main

import (
	"fmt"
	"runtime"
	"sync"
)

func gomaxProcs() {
	// กำหนดให้ใช้ CPU 2 Cores
	prev := runtime.GOMAXPROCS(2)
	fmt.Println("Previous GOMAXPROCS : ", prev)

	// อ่านค่า GOMAXPROCS ปัจจุบัน
	current := runtime.GOMAXPROCS(0)
	fmt.Println("Current GOMAXPROCS : ", current)
}

func cpuBound() {
	var wg sync.WaitGroup
	/*	หาก GOMAXPROCS = 1: Goroutines จะรันทีละ 1 บน 1 Core
		หาก GOMAXPROCS > 1: Goroutines จะรันพร้อมกันบนหลาย Core 
	*/
	current := runtime.GOMAXPROCS(2)
	fmt.Println("GOMAXPROCS : ", current)

	wg.Add(4)

	for i := 0; i < 4; i++ {
		go func(id int) {
			defer wg.Done()
			sum := 0
			for j := 0; j < 10; j++ {
				sum += j
			}
			fmt.Printf("Goroutine %d finished with sum: %d\n", id, sum)
		}(i)
	}
	wg.Wait()
}

func main() {
	// gomaxProcs()
	cpuBound()
}
