package main

import (
	"fmt"
	"runtime"
)

func allowcateMemory() {
	data := make([]byte, 10*1024*1024) // 10MB
	_ = data
}

func main() {
	var m runtime.MemStats
	allowcateMemory()
	runtime.ReadMemStats(&m)
	fmt.Printf("Before GC: Alloc = %v MiB\n", m.Alloc/1024/1024)

	runtime.GC() // target GC manually

	runtime.ReadMemStats(&m)
	fmt.Printf("After GC: Alloc = %v MiB\n", m.Alloc/1024/1024)

	fmt.Printf("NumGC = %v\n", m.NumGC) // จำนวนรอบที่ GC ทำงาน
	fmt.Printf("TotalAlloc = %v bytes\n", m.TotalAlloc)
}
