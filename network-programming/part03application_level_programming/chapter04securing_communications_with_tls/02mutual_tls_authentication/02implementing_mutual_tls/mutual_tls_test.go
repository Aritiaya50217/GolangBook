package test

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
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

func caCertPool(caCertFn string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(caCertFn)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to add certificate to pool")
	}

	return certPool, nil
}

func TestMutualTLSAuthentication(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverPool, err := caCertPool("clientCert.pem")
	if err != nil {
		t.Fatal(err)
	}

	cert, err := tls.LoadX509KeyPair("serverCert.pem", "serverKey.pem")
	if err != nil {
		t.Fatalf("loading key pair : %v", err)
	}

	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		GetConfigForClient: func(hello *tls.ClientHelloInfo) (*tls.Config, error) {
			return &tls.Config{
				Certificates:             []tls.Certificate{cert},
				ClientAuth:               tls.RequireAndVerifyClientCert,
				ClientCAs:                serverPool,
				CurvePreferences:         []tls.CurveID{tls.CurveP256},
				MinVersion:               tls.VersionTLS13,
				PreferServerCipherSuites: true,
				VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
					opts := x509.VerifyOptions{
						KeyUsages: []x509.ExtKeyUsage{
							x509.ExtKeyUsageClientAuth,
						},
						Roots: serverPool,
					}

					ip := strings.Split(hello.Conn.RemoteAddr().String(), ":")[0]
					hostnames, err := net.LookupAddr(ip)
					if err != nil {
						t.Errorf("PTR lookup: %v", err)
					}
					hostnames = append(hostnames, ip)

					for _, chain := range verifiedChains {
						opts.Intermediates = x509.NewCertPool()
						for _, cert := range chain[1:] {
							opts.Intermediates.AddCert(cert)
						}

						for _, hostname := range hostnames {
							opts.DNSName = hostname
							if _, err := chain[0].Verify(opts); err == nil {
								return nil
							}
						}
					}
					return errors.New("client authentication failed")
				},
			}, nil
		},
	}

	serverAddress := "localhost:44443"
	server := NewTLSServer(ctx, serverAddress, 0, serverConfig)
	done := make(chan struct{})

	go func() {
		if err := server.ListenAndServeTLS("serverCert.pem", "serverKey.pem"); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			t.Error(err)
			return
		}
		done <- struct{}{}
	}()
	server.Ready()

	clientPool, err := caCertPool("serverCert.pem")
	if err != nil {
		t.Fatal(err)
	}

	clientCert, err := tls.LoadX509KeyPair("clientCert.pem", "clientKey.pem")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := tls.Dial("tcp", serverAddress, &tls.Config{
		Certificates:     []tls.Certificate{clientCert},
		CurvePreferences: []tls.CurveID{tls.CurveP256},
		MinVersion:       tls.VersionTLS13,
		RootCAs:          clientPool,
	})

	if err != nil {
		t.Fatal(err)
	}

	hello := []byte("hello")
	if _, err = conn.Write(hello); err != nil {
		t.Fatal(err)
	}

	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err != nil {
		t.Fatal(err)
	}

	if actual := b[:n]; !bytes.Equal(hello, actual) {
		t.Fatal("expected %q; actual %q , hello , actual")
	}

	if err = conn.Close(); err != nil {
		t.Fatal(err)
	}
	cancel()

	<-done
}
