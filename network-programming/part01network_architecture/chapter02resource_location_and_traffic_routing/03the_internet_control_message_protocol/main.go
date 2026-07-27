package main

import (
	"fmt"
	"net"
)

func main() {
	ip, err := net.ResolveIPAddr("ip4", "8.8.8.8")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Resolved IP : ", ip)
}
