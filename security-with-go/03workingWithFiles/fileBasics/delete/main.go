package main

import (
	"log"
	"os"
)

func main() {
	filename := "/fileBasics/creatingAnEmptyFile/test.txt"
	err := os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
}
