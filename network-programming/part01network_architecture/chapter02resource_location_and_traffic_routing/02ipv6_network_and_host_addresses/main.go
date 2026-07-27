package main

import (
	"fmt"
	"net"
)

func main() {
	ip, network, err := net.ParseCIDR("2001:db8:abcd:1234::1/64")
	if err != nil {
		fmt.Println(err)
		return
	}

	ones, bits := network.Mask.Size()

	fmt.Println("IP Address : ", ip)
	fmt.Println("Network    :", network.IP)
	fmt.Println("Prefix Length  : ", ones)
	fmt.Println("Address    :", bits)
}
