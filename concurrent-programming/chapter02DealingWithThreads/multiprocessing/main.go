package main

import (
	"fmt"
	"runtime"
	"sync"
)

func compute(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[%s] running on core %d\n", name, runtime.GOMAXPROCS(0))
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Printf("[%s] done\n", name)
}
func main() {
	runtime.GOMAXPROCS(4) // ใช้ 4 core
	var wg sync.WaitGroup

	tasks := []string{"A", "B", "C", "D"}
	for _, t := range tasks {
		wg.Add(1)
		go compute(t, &wg)
	}
	wg.Wait()
	fmt.Println("All tasks completed!")
}
