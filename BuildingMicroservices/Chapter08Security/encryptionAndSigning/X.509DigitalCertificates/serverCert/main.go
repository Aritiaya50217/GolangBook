package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func writePEM(path string, block *pem.Block) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("create %s : %v", path, err)
	}
	defer f.Close()
	if err := pem.Encode(f, block); err != nil {
		log.Fatalf("write %s : %v", path, err)
	}
}

func main() {
	// CA key & cert
	caKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	caTmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Example CA Co"},
			CommonName:   "ExampleTestRootCA",
		},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
	}
	caDER, err := x509.CreateCertificate(rand.Reader, caTmpl, caTmpl, &caKey.PublicKey, caKey)
	if err != nil {
		log.Fatalf("create ca : %v", err)
	}
	writePEM("ca.pem", &pem.Block{Type: "CERTIFICATE", Bytes: caDER})
	writePEM("ca.key", &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caKey)})

	// parse ca cert for signing children
	caCert, err := x509.ParseCertificate(caDER)
	if err != nil {
		log.Fatalf("parse ca cert: %v", err)
	}

	// server key &cert signed by CA
	srvKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	srvTmpl := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			CommonName: "lacalhost",
		},
		NotBefore:   time.Now().Add(-time.Hour),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"localhost"},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}
	srvDER, err := x509.CreateCertificate(rand.Reader, srvTmpl, caCert, &srvKey.PublicKey, caKey)
	if err != nil {
		log.Fatalf("create sever cert: %v", err)
	}
	writePEM("server.crt", &pem.Block{Type: "CERTIFICATE", Bytes: srvDER})
	writePEM("server.key", &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(srvKey)})

	// client key & cert (for mTLS)
	clientKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	clientTmpl := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			CommonName: "test-client",
		},
		NotBefore: time.Now().Add(-time.Hour),
		NotAfter:  time.Now().AddDate(1, 0, 0),
		KeyUsage:  x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
		},
	}
	clientDER, err := x509.CreateCertificate(rand.Reader, clientTmpl, caCert, &clientKey.PublicKey, caKey)
	if err != nil {
		log.Fatalf("create client cert : %v", err)
	}

	writePEM("client.crt", &pem.Block{Type: "CERTIFICATE", Bytes: clientDER})
	writePEM("client.key", &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(clientKey)})

	log.Println("Wrote: ca.pem ca.key server.crt server.key client.crt client.key")
}
