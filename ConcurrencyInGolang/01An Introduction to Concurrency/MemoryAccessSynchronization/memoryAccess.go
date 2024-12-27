package main

import (
	"fmt"
	"sync"
)

func MemoryAccess() {
	var value int
	var memoryAccess sync.Mutex
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if value == 0 {
		fmt.Println("the value is ", value)
	} else {
		fmt.Println("the value is ", value)
	}
	memoryAccess.Unlock()
}

func main() {
	// var data int
	// go func() {
	// 	data++
	// }()
	// // time.Sleep(1 * time.Second)
	// if data == 0 {
	// 	fmt.Println("the value is 0 ")
	// } else {
	// 	fmt.Println("the value is ", data)
	// }
	MemoryAccess()
}
