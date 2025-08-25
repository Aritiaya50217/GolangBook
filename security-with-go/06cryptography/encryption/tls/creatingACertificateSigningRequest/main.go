package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"os"
)

func main() {
	// create private key
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// create CSR template
	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{"My Org"},
			CommonName:   "example.com",
		},
	}

	// create CSR
	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)

	// save CSR
	csrOut, _ := os.Create("csr.pem")
	pem.Encode(csrOut, &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	})
	csrOut.Close()

	println("CSR created (csr.pem)")
}
