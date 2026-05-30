package main

import (
	"fmt"
	"runtime/debug"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

/*
KB = 2^10
MB = 2^20
GB = 2^30
TB = 2^40
*/
func main() {
	fmt.Println("GOMEMLIMIT : ", debug.SetMemoryLimit(1<<30)) // 1 GB
}
