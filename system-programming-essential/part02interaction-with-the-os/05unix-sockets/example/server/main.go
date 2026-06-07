package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	socketPath := "/tmp/app.sock"

	// ลบ socket เก่าถ้ามี
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			buf := make([]byte, 1024)
			n, _ := c.Read(buf)

			fmt.Println("Received : ", string(buf[:n]))
			c.Write([]byte("pong"))
		}(conn)
	}
}
