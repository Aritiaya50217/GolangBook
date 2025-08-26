package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	// ----------------------
	// 1. สร้าง CA
	// ----------------------
	caKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"My Root CA"},
			CommonName:   "My Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caBytes, _ := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caKey.PublicKey, caKey)

	// บันทึก CA cert + key
	writePem("ca.pem", "CERTIFICATE", caBytes)
	writeECKey("ca.key", caKey)

	fmt.Println("CA created: ca.pem + ca.key")

	// ----------------------
	// 2. สร้าง Server Certificate
	// ----------------------
	serverKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	serverTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"MyOrg Server"},
			CommonName:   "localhost",
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"localhost"},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}

	serverBytes, _ := x509.CreateCertificate(rand.Reader, &serverTemplate, &caTemplate, &serverKey.PublicKey, caKey)
	writePem("cert.pem", "CERTIFICATE", serverBytes)
	writeECKey("key.pem", serverKey)

	fmt.Println("Server cert created: cert.pem + key.pem")
}

// ฟังก์ชันบันทึก PEM
func writePem(filename, blockType string, bytes []byte) {
	f, _ := os.Create(filename)
	defer f.Close()
	pem.Encode(f, &pem.Block{Type: blockType, Bytes: bytes})
}

// ฟังก์ชันบันทึก EC Key
func writeECKey(filename string, key *ecdsa.PrivateKey) {
	b, _ := x509.MarshalECPrivateKey(key)
	f, _ := os.Create(filename)
	defer f.Close()
	pem.Encode(f, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
}
