package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

// Process หนึ่งตัวมีหลาย goroutines (threads)
func runProcess(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting process %d...\n", id)

	// ตัวอย่างรัน process จริง
	cmd := exec.Command("sleep", "2")
	cmd.Start()

	// thread ภายใน process
	var innerWG sync.WaitGroup
	for j := 1; j <= 3; j++ {
		innerWG.Add(1)
		go func(pid, tid int) {
			defer innerWG.Done()
			fmt.Printf("Process %d - Thread %d working...\n", pid, tid)
			time.Sleep(time.Millisecond * 500)
		}(id, j)
	}
	innerWG.Wait()
	cmd.Wait()

	fmt.Printf("Process %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go runProcess(i, &wg)
	}
	wg.Wait()
	fmt.Println("All processes and threads done")
}
