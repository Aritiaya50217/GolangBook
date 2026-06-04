package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	log.Println("HTTPS running on :8443")

	log.Fatal(server.ListenAndServeTLS("", ""))
}

