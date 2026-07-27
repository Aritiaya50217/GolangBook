package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	start := time.Now()

	conn, err := net.DialTimeout("tcp", "google.com:80", 5*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}

	// วัด bandwidth
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	fmt.Println("Receive: ", n, " bytes")

	defer conn.Close()

	fmt.Println("Latency : ", time.Since(start))
}
