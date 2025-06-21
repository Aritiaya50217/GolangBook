package main

import (
	"fmt"
	"sync"
)

func main() {
	//  create a waitgroup
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		// make sure the waitgroup knows about
	// 		// goroutine completion
	// 		defer wg.Done()
	// 	}()
	// }
	// // wait until all goroutines are done
	// wg.Wait()

	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}
	// the is no goroutine reading from ch
	// none of the goroutines will return
	// so this will deadlock at Wait below
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	wg.Wait()
	close(ch)

}
