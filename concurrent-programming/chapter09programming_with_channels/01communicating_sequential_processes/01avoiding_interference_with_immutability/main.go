package main

import "fmt"

type Counter struct {
	value int
}

// ไม่แก้ของเดิม แต่สร้าง counter ใหม่
func increment(c Counter) Counter {
	return Counter{value: c.value + 1}
}

type Message struct {
	text string // ไม่แก้หลังส่ง
}

func main() {
	c1 := Counter{value: 0}
	c2 := increment(c1)
	c3 := increment(c1)

	fmt.Println("--- Avoiding Interference with Immutability ---")
	fmt.Println(c1.value)
	fmt.Println(c2.value)
	fmt.Println(c3.value)

	fmt.Println("--- Immutability + Channel (CSP Style) ---")
	ch := make(chan Message)
	go func() {
		msg := Message{text: "Hello"}
		ch <- msg
	}()

	received := <-ch
	fmt.Println(received.text)
}
