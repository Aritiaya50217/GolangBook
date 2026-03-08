package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
)

func Take[K any](quit chan int, n int, input <-chan K) <-chan K {
	output := make(chan K)
	go func() {
		defer close(output)
		moreData := true
		var msg K
		for n > 0 && moreData { // continues forwarding message as long as there is more data and countdown n is greater than 0
			select {
			case msg, moreData = <-input: // reads the next message from the input
				if moreData {
					output <- msg // forwards the message to the output
					n--           // reduces the countdown variable n by 1
				}
			case <-quit:
				return
			}
		}
		if n == 0 {
			close(quit) // closes the quit channel if the countdown reaches 0
		}
	}()
	return output
}

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

func Broadcast[K any](quit <-chan int, input <-chan K, n int) []chan K {
	outputs := CreateAll[K](n) // creates n output channels of type K (see the next listing for the implementation)
	go func() {
		defer CloseAll(outputs...) // once complete , closes all putput channels (see the next listing for the implementation)
		var msg K
		moreData := true
		for moreData {
			select {
			case msg, moreData = <-input: // reads the next message from the input channel
				if moreData {
					for _, output := range outputs { // if the input channel hasn't been closed, writes the message to each output channel
						output <- msg
					}
				}
			case <-quit:
				return
			}
		}
	}()
	return outputs
}

func CreateAll[K any](n int) []chan K { // creates n channels of type K
	channels := make([]chan K, n)
	for i, _ := range channels {
		channels[i] = make(chan K)
	}
	return channels

}

func CloseAll[K any](channels ...chan K) { // closes all the channels
	for _, output := range channels {
		close(output)
	}
}

func longestsWords(quit <-chan int, words <-chan string) <-chan string {
	longWords := make(chan string)
	go func() {
		defer close(longWords)
		uniqueWordsMap := make(map[string]bool) // creates a map to store unique words
		uniqueWords := make([]string, 0)        // creates slice to store the slice the list of unique words for easy sorting later
		moreData, word := true, ""
		for moreData {
			select {
			case word, moreData = <-words:
				if moreData && !uniqueWordsMap[word] { // if the channel is not closed and the word is a new one , adds the new word to the map and list
					uniqueWordsMap[word] = true
					uniqueWords = append(uniqueWords, word)
				}

			case <-quit:
				return
			}
		}
		sort.Slice(uniqueWords, func(i, j int) bool {
			return len(uniqueWords[i]) > len(uniqueWords[j])
		})
		longWords <- strings.Join(uniqueWords[:10], ",") // once the input channel is closed , sends a string with the 10 longest words on the output channel
	}()
	return longWords
}

func frequentWords(quit <-chan int, words <-chan string) <-chan string {
	mostFrequentWords := make(chan string)
	go func() {
		defer close(mostFrequentWords)
		freqMap := make(map[string]int) // creates a map to store the frequency occurrence of each unique word
		freqList := make([]string, 0)   // creates a slice to store a list of unique words
		moreData, word := true, ""
		for moreData {
			select {
			case word, moreData = <-words: // consumes the next message on the input channel
				if moreData {
					if freqMap[word] == 0 {
						freqList = append(freqList, word) // if the message contains a new word , adds it to the slice of unique words
					}
					freqMap[word] += 1 // increments the count of the word
				}
			case <-quit:
				return
			}
		}
		sort.Slice(freqList, func(i, j int) bool {
			return freqMap[freqList[i]] > freqMap[freqList[j]] // once all input messages are consumed, sorts the list of words by the occurrence count
		})
		mostFrequentWords <- strings.Join(freqList[:10], ", ") // writes the 10 most frequent words onto the output channel
	}()
	return mostFrequentWords
}

const downloaders = 20

func main() {
	quitWords := make(chan int) // creates a separate quit channel to be used before the Take(n) function
	quit := make(chan int)
	defer close(quit)
	urls := generateUrls(quitWords)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPages(quitWords, urls)
	}
	words := Take(quitWords, 10000, extractWords(quitWords, FanIn(quitWords, pages...))) // creates the Take(n) gorotine with a 10000 countdown , feeding from the extractWords() output
	wordsMulti := Broadcast(quit, words, 2)                                              // uses a separate quit channel for the rest of the pipeline
	longestResults := longestsWords(quit, wordsMulti[0])
	frequentResults := frequentWords(quit, wordsMulti[1])

	fmt.Println("Longest Words:", <-longestResults)        // Read he result from the longestWords() goroutine and prints is.
	fmt.Println("Most frequent Words:", <-frequentResults) // read the result from the mostFrequentWords() goroutine and prints it.

}
