package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
