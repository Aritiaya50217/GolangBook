package main

import "fmt"

func main() {
	var data [][]byte

	for i := 0; i < 100000; i++ {
		b := make([]byte, 1024)
		data = append(data, b)
	}
	fmt.Println("done")

	// GOGC=20 GODEBUG=gctrace=1 go run main.go  --> GC บ่อย ,RAM น้อย, CPU มาก
	// GOGC=10000 GODEBUG=gctrace=1 go run main.go -->  GC น้อย, RAM มาก, CPU น้อย
}
