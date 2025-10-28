package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
	resp, _ := http.Get(url) // downloads the web page from the given URL
	defer resp.Body.Close()  // close the response at the end of the function

	if resp.StatusCode != 200 {
		panic("Server returning error status code : " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body { // iterates over every downloaded character
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c) // finds the index of the character in the alphabet
		if cIndex >= 0 {
			frequency[cIndex] += 1 // if the character is part of the alphabet , increments the count by 1
		}
	}
	fmt.Println("Completed : ", url)
}

func main() {
	var frequency = make([]int, 26) // initialzes slice space for the frequency table
	for i := 1000; i <= 1030; i++ { // iterates from document ID 1000 to 1030 to download 31 docs
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency) // calls the countLetters() function sequentially
	}

	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i]) // Outputs each letter with its frequency
	}
}
