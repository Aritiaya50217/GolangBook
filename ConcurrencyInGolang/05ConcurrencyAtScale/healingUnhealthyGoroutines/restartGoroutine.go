package main

import (
	"fmt"
	"time"
)

func workerWithRestart(restartChan chan bool) {
	for {
		select {
		case <-restartChan:
			fmt.Println("Restarting worker... ")
			go workerWithRestart(restartChan)
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	restartChan := make(chan bool)
	go workerWithRestart(restartChan)

	time.Sleep(3 * time.Second)
	fmt.Println("Signaling worker to restart...")
	restartChan <- true // ส่งสัญญาณให้ Restart

	time.Sleep(2 * time.Second)
	fmt.Println("Main function exits.")

}
