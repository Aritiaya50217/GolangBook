package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	echoCmd := exec.Command("echo", "Hello, world!")
	grepCmd := exec.Command("grep", "Hello")

	pipe, err := echoCmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// ให้ grep อ่านจาก pipe
	grepCmd.Stdin = pipe

	// แสดงผล grep ออกหน้าจอ
	grepCmd.Stdout = os.Stdout

	if err := grepCmd.Start(); err != nil {
		panic(err)
	}

	if err := echoCmd.Run(); err != nil {
		panic(err)
	}

	if err := grepCmd.Wait(); err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
