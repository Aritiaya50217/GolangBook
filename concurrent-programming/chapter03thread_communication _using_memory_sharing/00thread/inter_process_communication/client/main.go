package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	conn.Write([]byte("Hello from Client"))
	reply := make([]byte, 1024)
	n, _ := conn.Read(reply)
	fmt.Println("Server says: ", string(reply[:n]))
	conn.Close()
}
