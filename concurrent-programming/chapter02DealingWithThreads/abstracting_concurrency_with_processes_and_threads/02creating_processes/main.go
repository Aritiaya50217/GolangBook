package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running %s: %v\n", name, err)
		return
	}
	fmt.Printf("[%s output]:\n%s\n", name, out)
}

func main() {
	var wg sync.WaitGroup
	commands := [][]string{
		{"echo", "Process 1 running"},
		{"echo", "Process 2 running"},
		{"echo", "Process 3 running"},
	}

	for _, cmd := range commands {
		wg.Add(1)
		go func(cmd []string) {
			defer wg.Done()
			runCommand(cmd[0], cmd[1:]...)
		}(cmd)
	}
	wg.Wait()
	fmt.Println("All processes finished")
}
