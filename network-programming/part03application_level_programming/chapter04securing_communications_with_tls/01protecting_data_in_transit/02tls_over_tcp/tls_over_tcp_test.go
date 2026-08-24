package tlsovertcp

import (
	"crypto/tls"
	"net"
	"testing"
	"time"
)

func TestClientTLSGoogle(t *testing.T) {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 30 * time.Second}, "tcp", "www.google.com:443",
		&tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	state := conn.ConnectionState()
	t.Logf("TLS 1.%d", state.Version-tls.VersionTLS10)
	t.Log(tls.CipherSuiteName(state.CipherSuite))
	t.Log(state.VerifiedChains[0][0].Issuer.Organization[0])
	_ = conn.Close()
}
