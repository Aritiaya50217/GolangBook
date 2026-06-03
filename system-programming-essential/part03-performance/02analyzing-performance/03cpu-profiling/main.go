package main

import (
	"os"
	"runtime/pprof"
)

func heavyWork() {
	sum := 0

	for i := 0; i < 1000000000; i++ {
		sum += i
	}

	_ = sum
}

func main() {
	f, _ := os.Create("cpu.prof")
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	heavyWork()
}
