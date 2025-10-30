package main

import (
	"fmt"
	"sync"
	"time"
)

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 10; i++ {
		mutex.Lock() // lock the mutex before entering the critical section
		*money += 10
		mutex.Unlock() // unlock after exiting the critical section
	}
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		*money -= 10
		mutex.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mutex := sync.Mutex{} // create a new mutex

	go stingy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Money in bank account : ", money)
	mutex.Unlock()
}
