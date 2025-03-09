package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var i int
	var v atomic.Value

	go func() {
		// This goroutine will eventually store 1 in v
		i = 1
		v.Store(1)
	}()

	go func() {
		// busy-waiting
		for {
			// This will keep checking until v has 1
			if val, _ := v.Load().(int); val == 1 {
				fmt.Println(i)
				return
			}
		}
	}()
	select {}
}
