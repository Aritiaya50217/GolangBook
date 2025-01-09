package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Error    error
	Response *http.Response
}

func checkStatus() {
	check := func(done <-chan interface{}, urls ...string) <-chan Result {
		responses := make(chan Result)
		go func() {
			defer close(responses)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{
					Error:    err,
					Response: resp,
				}
				select { // ใช้เลือก channel ที่พร้อมใช้งานใน goroutines
				case <-done:
					return
				case responses <- result:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range check(done, urls...) {
		if response.Error != nil {
			fmt.Printf("error: %v", response.Error)
			continue
		}
		fmt.Printf("Response: %v\n", response.Response.Status)
	}

}

func errCounts() {
	check := func(done <-chan interface{}, urls ...string) <-chan Result {
		responses := make(chan Result)
		go func() {
			defer close(responses)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{
					Error:    err,
					Response: resp,
				}
				select { // ใช้เลือก channel ที่พร้อมใช้งานใน goroutines
				case <-done:
					return
				case responses <- result:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	errCount := 0
	urls := []string{"a", "https://www.google.com", "b", "c", "d"}
	for result := range check(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

func worker(id int, wg *sync.WaitGroup, mu *sync.Mutex, errors *[]error) {
	defer wg.Done()
	if id == 3 {
		mu.Lock()
		*errors = append(*errors, fmt.Errorf("worker %d failed", id))
		mu.Unlock()
	}
}

func waitGroup() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &mu, &errs)
	}
	wg.Wait()

	if len(errs) > 0 {
		fmt.Println("Errors encountered:")
		for _, err := range errs {
			fmt.Println(err)
		}
	}
	fmt.Println("All workers completed successfully")
}

func worker2(ctx context.Context, id int, ch chan error) {
	select {
	case <-time.After(2 * time.Second):
		if id == 1 {
			ch <- errors.New("worker 1 failed")
		} else {
			ch <- nil
		}
	case <-ctx.Done():
		ch <- ctx.Err()
	}
}

func useContext() {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	numWorkers := 3
	errChannel := make(chan error, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker2(ctx, i, errChannel)
	}

	for i := 1; i <= numWorkers; i++ {
		if err := <-errChannel; err != nil {
			fmt.Println("Error : ", err)
		} else {
			fmt.Println("Worker completed successfully")
		}
	}

}

func main() {
	// checkStatus()
	// errCounts()
	// waitGroup()
	useContext()

}
