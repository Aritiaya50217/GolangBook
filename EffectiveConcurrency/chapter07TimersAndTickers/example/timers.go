package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("Waiting for 3 seconds...")

	<-timer.C // รอจนกว่าจะถึงเวลา
	fmt.Println("Timer expired!")
}
