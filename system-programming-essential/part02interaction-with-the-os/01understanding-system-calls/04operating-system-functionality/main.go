package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// start a new process
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// Get the current process ID
	pid := os.Getegid()
	fmt.Println("Current process ID : ", pid)
}
