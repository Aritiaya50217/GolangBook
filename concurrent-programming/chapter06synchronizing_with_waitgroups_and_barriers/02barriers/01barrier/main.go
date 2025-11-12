package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	const (
		numWorkers = 3
		numPhases  = 2
	)
	arrived := 0
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for phase := 1; phase <= numPhases; phase++ {
				fmt.Printf("Worker %d doing phase %d\n", id, phase)
				mu.Lock()
				arrived++
				if arrived == numWorkers {
					// ทุก worker มาถึง barrier แล้ว
					fmt.Printf("=== All workers reached barrier for phase %d ===\n", phase)
					arrived = 0
					cond.Broadcast() // ปล่อยทุกคนไปต่อ
				} else {
					cond.Wait() // รอคนอื่นให้ครบ
				}
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
}
