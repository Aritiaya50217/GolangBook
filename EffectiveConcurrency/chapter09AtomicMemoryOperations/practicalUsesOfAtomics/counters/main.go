package main

import (
	"fmt"
	"sync/atomic"
)

var count int64

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
		}()
	}

	for {
		v := atomic.LoadInt64(&count)
		fmt.Println(v)
		if v == 1 {
			break
		}
	}
}
