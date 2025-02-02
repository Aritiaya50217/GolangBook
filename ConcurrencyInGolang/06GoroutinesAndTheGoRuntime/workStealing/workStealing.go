package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 10

	// เริ่ม goroutine จำนวน 10 ตัว
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// รอให้ทุก goroutine ทำงานเสร็จ
	wg.Wait()

}
