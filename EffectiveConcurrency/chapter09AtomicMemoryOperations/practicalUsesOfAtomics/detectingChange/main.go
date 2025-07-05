package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var value int32 = 0
	swapped := atomic.CompareAndSwapInt32(&value, 0, 1)

	if swapped {
		fmt.Println("Value changed from 0 to 1")
	} else {
		fmt.Println("Value was not 0,so not changed")
	}
	fmt.Println("Current value:", atomic.LoadInt32(&value))
}
