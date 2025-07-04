package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var str atomic.Value
	var done atomic.Bool
	var a int
	str.Store("")
	go func() {
		a = 1
		str.Store("Done!")
		done.Store(true)
	}()
	for !done.Load() {
		fmt.Println(a)
	}
	fmt.Println(str.Load())
}
