package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Stat returns file info. It will return
	// an error if there is no file.
	filename := "/fileBasics/creatingAnEmptyFile/test.txt"
	fileInfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Fatalf("File %s does not exist", filename)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
