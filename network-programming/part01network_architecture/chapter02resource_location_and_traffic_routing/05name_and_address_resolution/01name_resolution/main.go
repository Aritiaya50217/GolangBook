package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupHost("google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ips)
}
