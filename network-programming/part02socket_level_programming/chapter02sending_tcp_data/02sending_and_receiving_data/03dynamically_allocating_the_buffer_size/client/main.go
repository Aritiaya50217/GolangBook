package main

import (
	"encoding/binary"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")

	defer conn.Close()

	data := []byte("Hello Golang")

	binary.Write(conn, binary.BigEndian, uint32(len(data)))

	conn.Write(data)
}
