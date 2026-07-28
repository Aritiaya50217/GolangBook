package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Connection failed : ", err)
		return
	}

	defer conn.Close()

	fmt.Println("Connected successfully")
}
