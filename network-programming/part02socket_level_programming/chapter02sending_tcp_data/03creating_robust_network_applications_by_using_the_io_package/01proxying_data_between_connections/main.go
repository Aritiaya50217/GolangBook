package main

import (
	"io"
	"log"
	"net"
)

func handle(client net.Conn) {
	defer client.Close()

	backend, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		return
	}

	defer backend.Close()

	go io.Copy(backend, client)
	io.Copy(client, backend)
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		client, _ := ln.Accept()

		go handle(client)
	}
}
