package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

func main() {
	pemFile := "../serverCert/ca.pem"
	certFile := "../serverCert/client.crt"
	keyFile := "../serverCert/client.key"

	caPEM, _ := os.ReadFile(pemFile)
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(caPEM)

	// สำหรับ mTLS ให้โหลด client cert:
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		// ถ้าไม่ใช้ mTLS ไม่ต้อง panic
		fmt.Println("no client cert for mTLS : ", err)
		panic(err)
	}

	tlsCfg := &tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{
		TLSClientConfig: tlsCfg,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:8443/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	n, _ := resp.Body.Read(buf)
	fmt.Println(string(buf[:n]))
}
