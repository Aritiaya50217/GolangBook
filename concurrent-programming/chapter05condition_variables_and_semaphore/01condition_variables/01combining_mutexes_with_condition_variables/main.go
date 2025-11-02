package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func stingy(money *int, cond *sync.Cond) {
	for i := 0; i <= 5; i++ {
		cond.L.Lock()
		fmt.Printf("index : %d, Money in stingy() : %d\n", i, *money)
		*money += 10
		cond.Signal() // signals on the condition variable every time we add to the shared money variable
		cond.L.Unlock()
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int, cond *sync.Cond) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("index : %d, Money in spendy() : %d\n", i, *money)
		cond.L.Lock()
		for *money < 10 {
			cond.Wait() // waits while we don't have enough money , releasing mutex and suspending execution
		}
		*money -= 10 // returning from wait() , reacquieres the mutex and subtracts money once there is enough money
		if *money < 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		cond.L.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex) // creates a new condition variable using a mutex
	// passes the condition variable to both goroutine
	go stingy(&money, cond)
	go spendy(&money, cond)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Money in bank account : ", money)
	mutex.Unlock()
}
