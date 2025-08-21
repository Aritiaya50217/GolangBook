package main

import (
	"log"
	"os"
)

func main() {
	// simple read only open. We will cover actually reading
	// and writing to files in examples further down the page
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// open file with more options. Last param is the permission mode
	// second param is the attributes when opening
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666) // 0666 คือ user/group/other สามารถอ่านและเขียนได้
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

}
