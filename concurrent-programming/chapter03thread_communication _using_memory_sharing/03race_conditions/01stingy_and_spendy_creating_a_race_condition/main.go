package main

import (
	"fmt"
	"time"
)

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	// Starts goroutines and passes a reference to the money variable
	go stingy(&money)
	go spendy(&money)

	time.Sleep(2 * time.Second)
	fmt.Println("Money in bank account : ", money)

}
