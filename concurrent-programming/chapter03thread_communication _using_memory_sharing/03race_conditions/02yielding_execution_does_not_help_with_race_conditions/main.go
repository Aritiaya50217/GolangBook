package main

import (
	"fmt"
	"runtime"
	"time"
)

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
		runtime.Gosched() // calls the Go scheduler after we perform the addition
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
		runtime.Gosched() // calls the Go scheduler after we perform the subtraction
	}
	fmt.Println("Spendy Done")
}
func main() {
	money := 100
	go stingy(&money)
	go spendy(&money)
	time.Sleep(2 * time.Second)
	fmt.Println("Money in bank account : ", money)
}
