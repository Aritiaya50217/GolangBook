package main

import "net"

func main() {
	conn, _ := net.Dial("udp", ":8080")

	conn.Write([]byte("hello"))
}
