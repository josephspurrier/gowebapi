package server_test

import (
	"app/webapi/pkg/server"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockServer struct {
	startHTTPCalled bool
	startHTTPError  error

	startHTTPSCalled bool
	startHTTPSError  error
}

func (s *MockServer) ListenAndServe() error {
	s.startHTTPCalled = true
	return s.startHTTPError
}

func (s *MockServer) ListenAndServeTLS(certFile, keyFile string) error {
	s.startHTTPSCalled = true
	return s.startHTTPSError
}

type MockLogger struct{}

func (l *MockLogger) Fatalf(format string, v ...interface{}) {}
func (l *MockLogger) Printf(format string, v ...interface{}) {}

func TestNoStart(t *testing.T) {
	s := new(server.Config)
	m := new(MockServer)
	ms := new(MockServer)
	l := new(MockLogger)

	s.Run(m, ms, l)

	assert.Equal(t, false, m.startHTTPCalled)
	assert.Equal(t, false, m.startHTTPSCalled)
	assert.Equal(t, false, ms.startHTTPCalled)
	assert.Equal(t, false, ms.startHTTPSCalled)
}

func TestHTTPStart(t *testing.T) {
	s := new(server.Config)
	m := new(MockServer)
	ms := new(MockServer)
	l := new(MockLogger)

	s.UseHTTP = true

	s.Run(m, ms, l)

	assert.Equal(t, true, m.startHTTPCalled)
	assert.Equal(t, false, m.startHTTPSCalled)
	assert.Equal(t, false, ms.startHTTPCalled)
	assert.Equal(t, false, ms.startHTTPSCalled)
}

func TestHTTPSStart(t *testing.T) {
	s := new(server.Config)
	m := new(MockServer)
	ms := new(MockServer)
	l := new(MockLogger)

	s.UseHTTPS = true

	s.Run(m, ms, l)

	assert.Equal(t, false, m.startHTTPCalled)
	assert.Equal(t, false, m.startHTTPSCalled)
	assert.Equal(t, false, ms.startHTTPCalled)
	assert.Equal(t, true, ms.startHTTPSCalled)
}

func TestBothStart(t *testing.T) {
	s := new(server.Config)
	m := new(MockServer)
	ms := new(MockServer)
	l := new(MockLogger)

	s.UseHTTP = true
	s.UseHTTPS = true

	s.Run(m, ms, l)

	// Sleep to ensure the Go routine runs for HTTPS.
	time.Sleep(2 * time.Second)

	assert.Equal(t, true, m.startHTTPCalled)
	assert.Equal(t, false, m.startHTTPSCalled)
	assert.Equal(t, false, ms.startHTTPCalled)
	assert.Equal(t, true, ms.startHTTPSCalled)
}

func TestHTTPAddress(t *testing.T) {
	s := new(server.Config)
	assert.Equal(t, ":80", s.HTTPAddress())
	s.HTTPPort = 8080
	assert.Equal(t, ":8080", s.HTTPAddress())
	s.Hostname = "example.com"
	assert.Equal(t, "example.com:8080", s.HTTPAddress())
}

func TestHTTPSAddress(t *testing.T) {
	s := new(server.Config)
	assert.Equal(t, ":443", s.HTTPSAddress())
	s.HTTPSPort = 4433
	assert.Equal(t, ":4433", s.HTTPSAddress())
	s.Hostname = "example.com"
	assert.Equal(t, "example.com:4433", s.HTTPSAddress())
}
