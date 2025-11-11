package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileSearch(dir, filename string, wg *sync.WaitGroup) {
	files, _ := os.ReadDir(dir) // reads all files from the directory given to the function
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name()) // joins each file to the directory
		if strings.Contains(file.Name(), filename) {
			fmt.Println(fpath) // if there is a match,print path on console.
		}

		if file.IsDir() {
			wg.Add(1)                          // if it is a directory, adds 1 to the waitgroup before starting a new goroutine.
			go fileSearch(fpath, filename, wg) // creates goroutine recursively, searching in the new directory.
		}
	}
	wg.Done() // marks Done() on the waitgroup after processing all files.
}

func main() {
	wg := sync.WaitGroup{}                     // creates a new , empty waitgroup
	wg.Add(1)                                  // adds a delta of 1 to the waitgroup.
	go fileSearch(os.Args[1], os.Args[2], &wg) // creates a new goroutine , performing the file serch and passing a reference to the waitgroup.
	wg.Wait() // waits for the search to complete. 
}
