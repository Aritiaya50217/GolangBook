package main

import "fmt"

func orDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func main() {
	bride := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				var stream <-chan interface{}
				select {
				case mybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = mybeStream
				case <-done:
					return
				}

				for val := range orDone(done, stream) {
					select {
					case valueStream <- val:
					case <-done:

					}
				}
			}
		}()
		return valueStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	done := make(chan interface{})
	defer close(done)

	for v := range bride(done, genVals()) {
		fmt.Printf("%v ", v)
	}
}
