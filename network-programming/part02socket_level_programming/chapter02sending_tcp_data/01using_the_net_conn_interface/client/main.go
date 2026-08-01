package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	conn.Write([]byte("Hello"))

	buf := make([]byte, 1024)

	n, _ := conn.Read(buf)

	fmt.Println(string(buf[:n]))
}
