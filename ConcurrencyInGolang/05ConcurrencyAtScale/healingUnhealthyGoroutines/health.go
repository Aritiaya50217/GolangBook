package main

import (
	"fmt"
	"time"
)

func monitorHealth(healthChan <-chan string, doneChan chan<- bool) {
	for {
		select {
		case msg := <-healthChan:
			fmt.Println("Health check:", msg)
		case <-time.After(1 * time.Second):
			fmt.Println("Goroutine might be unhealthy!")
			doneChan <- true
			return
		}
	}
}

func worker(healthChan chan<- string) {
	for {
		healthChan <- "alive"
		time.Sleep(1 * time.Second)
	}
}
func main() {

	healthChan := make(chan string)
	doneChan := make(chan bool)

	go monitorHealth(healthChan, doneChan)
	go worker(healthChan)

	<-doneChan // รอจนกระทั่งระบบตรวจสอบว่ามีปัญหา
	fmt.Println("Main function exits.")

}
