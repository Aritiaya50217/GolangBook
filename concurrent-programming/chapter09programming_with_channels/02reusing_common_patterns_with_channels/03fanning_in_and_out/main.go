package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func generateUrls(quit <-chan int) <-chan string { // accepts the quit channel and returns the output channel
	urls := make(chan string) // creates the output channel
	go func() {
		defer close(urls) // once complete, closes the output channel
		for i := 100; i <= 130; i++ {
			url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
			select {
			case urls <- url: // writes 50 urls to the output channel
			case <-quit:
				return

			}
		}
	}()
	return urls // returns the output channel
}

func downloadPages(quit <-chan int, urls <-chan string) <-chan string {
	pages := make(chan string) // creates the output channel, which will contrain downloaded web pages
	go func() {
		defer close(pages) // closes the output channel when it's finished
		moreData, url := true, ""
		for moreData { // continu
			select {
			case url, moreData = <-urls: // updates variables with a new message and flag to show whether there is more data
				if moreData {
					resp, _ := http.Get(url)
					if resp.StatusCode != 200 {
						panic("Server's error : " + resp.Status)
					}
					body, _ := io.ReadAll(resp.Body)
					pages <- string(body)
					resp.Body.Close()
				}
			case <-quit: // when a message arrives on the quit channel , terminates the goroutine
				return
			}
		}
	}()
	return pages // returns the output channel
}

func extractWords(quit <-chan int, pages <-chan string) <-chan string {
	words := make(chan string) // cretes the output channel , which will contain extracted words
	go func() {
		defer close(words)
		wordRegex := regexp.MustCompile(`[a-zA-Z]+`) // creates a regular expression to extract the word
		moreData, pg := true, ""
		for moreData {
			select {
			case pg, moreData = <-pages: // updates  variables with a new message and flag to show whether there is more data
				if moreData {
					for _, word := range wordRegex.FindAllString(pg, -1) {
						words <- strings.ToLower(word) // when a new text page is received , extracts all word with the regex and sends them on the output channel
					}
				}
			case <-quit: // when a message arrives on the quit channel , terminates goroutine
				return
			}
		}
	}()
	return words // return the output channel
}

func FanIn[K any](quit <-chan int, allChannels ...<-chan K) chan K {
	wg := sync.WaitGroup{}
	wg.Add(len(allChannels)) // Creates. a waitgroup , setting the size to be equal to the number of input channels
	output := make(chan K)   // Creates the output channel

	for _, c := range allChannels {
		go func(channel <-chan K) { // Starts a goroutine for every input channel
			defer wg.Done() // once the goroutine terminates , marks the waitgroup as done
			for i := range channel {
				select {
				case output <- i: // Forwards each received message to the shared output channel
				case <-quit: // if quit channel is closed , terminates the goroutine
					return
				}
			}
		}(c) // Passes one input channel to the goroutine
	}

	go func() {
		wg.Wait() // waits for all the goroutines to finish and then closes the output channel
		close(output)
	}()
	return output // returns the output channel
}

const downloaders = 20

func main() {
	quit := make(chan int)
	defer close(quit)
	urls := generateUrls(quit)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPages(quit, urls)
	}
	results := extractWords(quit, FanIn(quit, pages...)) // joins all pages channels into one channel using the fan-in pattern
	for result := range results {
		fmt.Println(result)
	}
}
