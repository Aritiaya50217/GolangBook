package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string) <-chan []int {
	result := make(chan []int) // creates output channel of type int slice
	go func() {
		defer close(result)
		frequency := make([]int, 26) // creates a local frequency slice
		resp, err := http.Get(url)
		if err != nil {
			panic(errors.New(err.Error()))
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			panic("Server returning error code : " + resp.Status)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(errors.New(err.Error()))
		}

		for _, b := range body {
			c := strings.ToLower(string(b))
			cIndex := strings.Index(allLetters, c)
			if cIndex >= 0 {
				frequency[cIndex] += 1 // updates each character count in the local frequency slice
			}
		}
		fmt.Println("Completed : ", url)
		result <- frequency // once it's finished , the frequency slice is sent over the channel
	}()
	return result
}

func main() {
	results := make([]<-chan []int, 0)  // creates a slice to contain all output channles
	totalFrequencies := make([]int, 26) // creates a slice to store the frequency for each letter in the english alphabet
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		results = append(results, countLetters(url)) // creates a goroutine for each web page and stores the output channel in the results slice
	}

	for _, c := range results { // Iterates over each output channel
		frequencyResult := <-c // receives a message from each output channel containing the frequencies for one web page
		for i := 0; i < 26; i++ {
			totalFrequencies[i] += frequencyResult[i] // adds the frequency counts to the total frequencies for each letter
		}
	}

	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, totalFrequencies[i])
	}
}
