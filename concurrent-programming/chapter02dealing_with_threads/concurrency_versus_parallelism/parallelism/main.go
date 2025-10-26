package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4) // CPU 4 core
	var wg sync.WaitGroup
	wg.Add(4)

	for i := 0; i <= 4; i++ {
		go func(id int) {
			defer wg.Done()
			sum := 0
			for j := 0; j < 100000000; j++ {
				sum += j
			}
			fmt.Printf("Goroutine %d done, sum=%d\n", id, sum)
		}(i)
	}
	wg.Wait()
	fmt.Println("Parallel processing complete!")
}
