package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")

	defer conn.Close()

	reader := bufio.NewScanner(conn)

	fmt.Fprintln(conn, "Hello")
	fmt.Fprintln(conn, "World")
	fmt.Fprintln(conn, "Golang")

	for i := 0; i < 3; i++ {
		reader.Scan()

		fmt.Println(reader.Text())
	}
}
