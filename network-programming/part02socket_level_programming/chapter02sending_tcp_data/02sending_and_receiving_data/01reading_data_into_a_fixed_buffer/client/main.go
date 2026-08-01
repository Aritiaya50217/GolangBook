package main

import "net"

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")

	defer conn.Close()

	conn.Write([]byte("Hello TCP"))
}
