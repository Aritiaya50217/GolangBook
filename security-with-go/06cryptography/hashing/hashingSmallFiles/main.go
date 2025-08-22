package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage : " + os.Args[0] + " <filepath>")
	fmt.Println("Example: " + os.Args[0] + " document.txt")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1) // ออกจากโปรแกรมทันที โดยไม่รัน defer
	}
	return os.Args[1]
}

func main() {
	filename := checkArgs()

	//  get bytes from file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// hash the file and output results
	fmt.Printf("MD5 : %x\n", md5.Sum(data))
	fmt.Printf("SHA1 : %x\n", sha1.Sum(data))
	fmt.Printf("SHA256 : %x\n", sha256.Sum256(data))
	fmt.Printf("SHA512 : %x\n", sha512.Sum512(data))
}
