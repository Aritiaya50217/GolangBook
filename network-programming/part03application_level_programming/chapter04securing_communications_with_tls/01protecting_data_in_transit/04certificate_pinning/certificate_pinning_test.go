package test

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"testing"
	"time"
)

type Server struct {
	ctx   context.Context
	ready chan struct{}

	addr      string
	maxIdle   time.Duration
	tlsConfig *tls.Config
}

func NewTLSServer(ctx context.Context, address string, maxIdle time.Duration, tlsConfig *tls.Config) *Server {
	return &Server{
		ctx:       ctx,
		ready:     make(chan struct{}),
		addr:      address,
		maxIdle:   maxIdle,
		tlsConfig: tlsConfig,
	}
}

func (s *Server) ListenAndServeTLS(certFn, keyFn string) error {
	if s.addr == "" {
		s.addr = "localhost:443"
	}

	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("binding to tcp %s: %w", s.addr, err)
	}

	if s.ctx != nil {
		go func() {
			<-s.ctx.Done()
			_ = l.Close()
		}()
	}

	return s.ServeTLS(l, certFn, keyFn)
}

func (s Server) ServeTLS(l net.Listener, certFn, keyFn string) error {
	if s.tlsConfig == nil {
		s.tlsConfig = &tls.Config{
			CurvePreferences:         []tls.CurveID{tls.CurveP256},
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
		}
	}

	if len(s.tlsConfig.Certificates) == 0 && s.tlsConfig.GetCertificate == nil {
		cert, err := tls.LoadX509KeyPair(certFn, keyFn)
		if err != nil {
			return fmt.Errorf("loading key pair: %v", err)
		}

		s.tlsConfig.Certificates = []tls.Certificate{cert}
	}

	tlsListener := tls.NewListener(l, s.tlsConfig)
	if s.ready != nil {
		close(s.ready)
	}

	for {
		conn, err := tlsListener.Accept()
		if err != nil {
			return fmt.Errorf("accept : %v", err)
		}

		go func() {

			defer func() { _ = conn.Close() }()

			for {
				if s.maxIdle > 0 {
					if err := conn.SetDeadline(time.Now().Add(s.maxIdle)); err != nil {
						return
					}
				}

				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					return
				}

				if _, err := conn.Write(buf[:n]); err != nil {
					return
				}
			}
		}()
	}
}

func (s *Server) Ready() {
	if s.ready != nil {
		<-s.ready
	}
}

func TestEchoServerTLS(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverAddress := "localhost:34443"
	maxIdle := time.Second
	server := NewTLSServer(ctx, serverAddress, maxIdle, nil)
	done := make(chan struct{})

	go func() {
		if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			t.Error(err)
			return
		}

		done <- struct{}{}
	}()

	server.Ready()

	cert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		t.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		t.Fatal("failed to append certificate to pool")
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.CurveP256},
		MinVersion:       tls.VersionTLS12,
		RootCAs:          certPool,
	}

	conn, err := tls.Dial("tcp", serverAddress, tlsConfig)
	if err != nil {
		t.Fatal(err)
	}

	hello := []byte("hello")
	_, err = conn.Write(hello)
	if err != nil {
		t.Fatal(err)
	}

	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err != nil {
		t.Fatal(err)
	}

	if actual := b[:n]; !bytes.Equal(hello, actual) {
		t.Fatalf("expected %q; actual %q", hello, actual)
	}

	time.Sleep(2 * maxIdle)
	_, err = conn.Read(b)
	if err != io.EOF {
		t.Fatal(err)
	}

	err = conn.Close()
	if err != nil {
		t.Fatal(err)
	}

	cancel()
	<-done
}
