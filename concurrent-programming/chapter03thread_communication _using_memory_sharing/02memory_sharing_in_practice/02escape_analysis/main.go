package main

import "fmt"

func shareVariable() *int {
	x := 10
	return &x // Compiler จะย้าย x ไปอยู่ใน heap แทน เพราะต้องใช้ต่อภายนอก
}

func main() {
	p := shareVariable()
	fmt.Println(*p)
	
	// command : go run -gcflags="-m" main.go
}
