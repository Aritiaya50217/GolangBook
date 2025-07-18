package main

import (
	"fmt"
	"time"
)

type Event struct {
	ID   int
	Data string
}

// Producer สร้าง event แล้วส่งเข้า channel
func Producer(events chan<- Event) {
	for i := 1; i <= 5; i++ {
		e := Event{
			ID:   1,
			Data: fmt.Sprintf("Event #%d", i),
		}
		fmt.Println("Producing: ", e)
		events <- e
		time.Sleep(500 * time.Millisecond)
	}
	close(events)
}

// Consumer รับ event จาก channel แล้วประมวลผล
func Consumer(events <-chan Event) {
	for e := range events {
		fmt.Println("Consuming : ", e)
		// ประมวลผล event เช่น ส่งอีเมล, บันทึกฐานข้อมูล
		time.Sleep(1 * time.Second)

	}
}

func main() {
	events := make(chan Event)
	go Producer(events)
	Consumer(events)

}
