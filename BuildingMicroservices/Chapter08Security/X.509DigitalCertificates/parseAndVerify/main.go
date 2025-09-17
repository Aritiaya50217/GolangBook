package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	caFile := "../serverCert/ca.pem"
	caPEM, _ := os.ReadFile(caFile)
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(caPEM)

	certFile := "../serverCert/server.crt"
	crtPEM, _ := os.ReadFile(certFile)
	block, _ := pem.Decode(crtPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		panic("no cert")
	}
	cert, _ := x509.ParseCertificate(block.Bytes)
	opts := x509.VerifyOptions{
		Roots:   roots,
		DNSName: "localhost", // ตรวจ SAN/DNS
	}
	if chains, err := cert.Verify(opts); err != nil {
		fmt.Println("verify failed: ", err)
	} else {
		fmt.Println("verified chains : ", chains)
	}
}
