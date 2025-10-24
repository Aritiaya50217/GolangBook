package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func heavyWork(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("sha256sum", "/bin/bash") // ตัวอย่างงาน CPU-bound
	out, _ := cmd.Output()
	fmt.Printf("Process %d result: %s\n", id, out)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go heavyWork(i, &wg)
	}
	wg.Wait()
	fmt.Println("Multiprocessing finished")
}
