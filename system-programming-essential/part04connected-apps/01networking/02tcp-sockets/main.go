package main

import (
	"fmt"
	"net"
)

func main() {
	// start listening for connections
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	// accept connections in a loop
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("write error : ", err.Error())
		}
	}
}
