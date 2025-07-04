package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count int32 = 0
	old := atomic.LoadInt32(&count)
	new := old + 1
	// พยายามเปลี่ยนจาก 0 เป็น 1
	swapped := atomic.CompareAndSwapInt32(&count, 0, new)
	for {
		if swapped {
			break
		}
	}
	fmt.Println("Swapped ? ", swapped)
	fmt.Println("Current count: ", count)

}
