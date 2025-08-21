package main

import (
	"log"
	"os"
)

func main() {
	// stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File dose not exist.")
		}
	}
	log.Panicln("File dose exist. File infomation: ")
	log.Println(fileInfo)
}
