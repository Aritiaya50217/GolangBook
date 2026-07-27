package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("2001:db8::1")

	if ip.To4() == nil {
		fmt.Println("This is IPv6")
	}
}
