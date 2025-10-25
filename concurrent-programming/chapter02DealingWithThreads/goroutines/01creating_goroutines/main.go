package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("Work %d started at %s\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func printMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, "iteration", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	// main() ก็เป็น goroutine ตัวหนึ่งเช่นกัน
	fmt.Println("== doWork ==")
	for i := 0; i < 5; i++ {
		go doWork(i)
	}
	time.Sleep(2 * time.Second)

	fmt.Println("\n== printMessage ==")
	go printMessage("Goroutine 1")
	go printMessage("Goroutine 2")

	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine finished")

}
