package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// load CA ให้ server สามารถ verify client (สำหรับ mTLS)
	caPEM, _ := os.ReadFile("ca.pem")
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caPEM)

	tlsCfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientCAs:  certPool,
		// ClientAuth: tls.RequireAndVerifyClientCert, // เปลี่ยนเป็น NoClientCert ถ้าไม่ต้องการ mTLS
		ClientAuth: tls.NoClientCert,
	}

	srv := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsCfg,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Hello TLS")
		}),
	}
	log.Println("listening : 8443")
	certFile := "../serverCert/server.crt"
	keyFile := "../serverCert/server.key"
	log.Fatal(srv.ListenAndServeTLS(certFile, keyFile))
}
