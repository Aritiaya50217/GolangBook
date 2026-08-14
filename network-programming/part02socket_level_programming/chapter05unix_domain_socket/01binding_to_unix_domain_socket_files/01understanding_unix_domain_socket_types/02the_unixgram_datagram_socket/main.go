package main

import (
	"context"
	"net"
	"os"
)

func datagramEchoServer(ctx context.Context, network, addr string) (net.Addr, error) {
	s, err := net.ListenPacket(network, addr)
	if err != nil {
		return nil, err
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = s.Close()
			if network == "unixgram" {
				_ = os.Remove(addr)
			}
		}()

		buf := make([]byte, 1024)
		for {
			n, clientAddr, err := s.ReadFrom(buf)
			if err != nil {
				return
			}

			if _, err = s.WriteTo(buf[:n], clientAddr); err != nil {
				return
			}
		}
	}()

	return s.LocalAddr(), nil
}
