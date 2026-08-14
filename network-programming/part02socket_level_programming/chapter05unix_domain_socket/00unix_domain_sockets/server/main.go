package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	os.Remove("/tmp/echo.sock")

	ln, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Listening...")

	conn, _ := ln.Accept()
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	fmt.Println("Received : ", string(buf[:n]))
}
