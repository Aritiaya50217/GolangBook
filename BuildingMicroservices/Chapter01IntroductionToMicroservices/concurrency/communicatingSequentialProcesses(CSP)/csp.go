package main

import (
	"fmt"
	"time"
)

func say(msg string, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- msg // ส่งข้อความผ่าน channel
}

func main() {
	ch := make(chan string)
	go say("Hello", ch)
	fmt.Println(<-ch) // ส่งข้อความที่รับมา
}
