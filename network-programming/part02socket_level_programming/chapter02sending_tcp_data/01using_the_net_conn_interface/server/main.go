package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client : ", conn.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Disconnected")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Receive : ", string(buf[:n]))

		conn.Write(buf[:])
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
