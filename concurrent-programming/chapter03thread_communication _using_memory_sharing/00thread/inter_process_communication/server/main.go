package main

import (
	"fmt"
	"net"
)

func main() {
	// TCP Socket (Process A <-> Process B)
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()
	fmt.Println("Server listening...")

	for {
		conn, _ := ln.Accept()
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println("Received : ", string(buf[:n]))
		conn.Write([]byte("Message received"))
		conn.Close()
	}
}
