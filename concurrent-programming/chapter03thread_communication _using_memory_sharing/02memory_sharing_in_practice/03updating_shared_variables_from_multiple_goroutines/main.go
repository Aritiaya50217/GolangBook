package main

import "fmt"

func main() {
	countChan := make(chan int)
	done := make(chan bool)

	go func() {
		count := 0
		for delta := range countChan {
			count += delta
		}
		fmt.Println("Final Count : ", count)
		done <- true
	}()

	for i := 0; i < 5; i++ {
		countChan <- 1
	}
	close(countChan)
	<-done
}
