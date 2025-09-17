package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Config TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
