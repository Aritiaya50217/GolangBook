package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	base := context.Background()
	c1, cancel1 := context.WithTimeout(base, 1*time.Second)
	defer cancel1()

	c2, cancel2 := context.WithTimeout(base, 2*time.Second)
	defer cancel2()

	c3, cancel3 := context.WithTimeout(base, 3*time.Second)
	defer cancel3()

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		<-c1.Done()
		fmt.Println("c1 done")
	}()

	go func() {
		defer wg.Done()
		<-c2.Done()
		fmt.Println("c2 done")
	}()

	go func() {
		defer wg.Done()
		<-c3.Done()
		fmt.Println("c3 done")
	}()
	wg.Wait()

}
