package main

import "fmt"

func main() {
	doWork := func(words <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range words {
				fmt.Println("s : ", s)
			}
		}()
		return completed
	}
	doWork(nil)
	fmt.Println("Done.")

}
