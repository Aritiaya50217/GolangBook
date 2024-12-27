package main

import (
	"fmt"
	"sync"
)

func sayHello() {
	fmt.Println("form func sayHello")
}

func main() {
	go func() {
		fmt.Println("form func inline")
	}()

	go sayHello()

	var wg sync.WaitGroup
	sayhello := func() {
		defer wg.Done()
		fmt.Println("form use defer")
	}

	wg.Add(1)
	go sayhello()
	wg.Wait()

	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println("salutation : ", salutation)

	arr := []string{"hello", "greeting", "good day"}
	for _, val := range arr {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			fmt.Println(val)
		}(val)
	}
	wg.Wait()
}
