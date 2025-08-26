package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Secure World with TLS!")
}

func main() {
	// load TLS certificate
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		panic(err)
	}

	// TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", helloHandler)

	fmt.Println("TLS server running at https://localhost:8443")
	err = server.ListenAndServeTLS("", "") // ไม่ต้องใส่ cert/key อีกเพราะ TLSConfig มีแล้ว
	if err != nil {
		panic(err)
	}
}
