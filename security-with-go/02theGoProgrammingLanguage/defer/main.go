package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		log.Fatal("Error creating file.")
	}
	defer file.Close()

}
