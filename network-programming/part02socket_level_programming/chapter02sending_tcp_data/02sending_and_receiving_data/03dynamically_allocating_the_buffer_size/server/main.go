package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		var length uint32

		if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
			return
		}

		buf := make([]byte, length)

		if _, err := io.ReadFull(conn, buf); err != nil {
			return
		}

		fmt.Println(string(buf))

	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
