package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func main() {
	// create CA private key
	caPrivate, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"My CA Org"},
			CommonName:   "MyCA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	// create CA certificate
	caBytes, _ := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caPrivate.PublicKey, caPrivate)
	caCert, _ := x509.ParseCertificate(caBytes)

	// create client provate key + CSR
	clientPrivate, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{"Client Org"},
			CommonName:   "client.examle.coom",
		},
	}
	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, clientPrivate)
	csr, _ := x509.ParseCertificateRequest(csrBytes)

	// เซ็น CSR -> client certificate
	clientCertTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      csr.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
	certBytes, _ := x509.CreateCertificate(rand.Reader, &clientCertTemplate, caCert, csr.PublicKey, caPrivate)

	// บันทึก certificate
	certOut, _ := os.Create("client_cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	certOut.Close()
	println("Signed certificate created (client_cert.pem)")

}
