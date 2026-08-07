package main

import (
	"fmt"
	"net"
)

func main() {
	addr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 9000,
	}

	conn, _ := net.DialTCP("tcp", nil, addr)
	defer conn.Close()

	conn.Write([]byte("Chunk1\n"))
	conn.Write([]byte("Chunk2\n"))
	conn.Write([]byte("Chunk3\n"))

	conn.CloseWrite() // บอกว่าส่งครบแล้ว

	buf := make([]byte, 1024)

	n, _ := conn.Read(buf)

	fmt.Println(string(buf[:n]))
}
