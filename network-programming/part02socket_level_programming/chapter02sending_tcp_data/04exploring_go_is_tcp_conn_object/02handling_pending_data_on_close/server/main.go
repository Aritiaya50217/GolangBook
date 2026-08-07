package main

import (
	"fmt"
	"io"
	"net"
)

// Half Close (CloseWrite)
func main() {
	ln, _ := net.ListenTCP("tcp", &net.TCPAddr{Port: 9000})
	defer ln.Close()

	conn, _ := ln.AcceptTCP()
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("Upload Complete")
			break
		}

		fmt.Println(string(buf[:n]))
	}

	conn.Write([]byte("Upload Success"))
}
