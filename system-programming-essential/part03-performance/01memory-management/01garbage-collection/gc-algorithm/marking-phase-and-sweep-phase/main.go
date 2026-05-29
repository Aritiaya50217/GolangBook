package main

import "fmt"

type User struct {
	Name string
}

func main() {
	u1 := &User{Name: "Alice"}
	u2 := &User{Name: "Bob"}

	_ = u1
	u2 = nil

	fmt.Println("mark phase : ", u1)
	fmt.Println("sweep phase : ", u2) // เป็น nil เพราะไม่มี object ให้ mark

}
