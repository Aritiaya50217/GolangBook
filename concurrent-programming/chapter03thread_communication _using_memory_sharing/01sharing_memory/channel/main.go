package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10 // ส่งข้อมูล
	}()
	value := <-ch // รับข้อมูล
	fmt.Println("Got : ", value)
}
