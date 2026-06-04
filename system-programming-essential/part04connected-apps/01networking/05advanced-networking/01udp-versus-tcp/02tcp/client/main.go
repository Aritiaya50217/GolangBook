package main

import "net"

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	conn.Write([]byte("Hello TCP"))
}
