package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		msg := scanner.Text()

		fmt.Println("Receive : ", msg)

		conn.Write([]byte("Echo : " + msg + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
