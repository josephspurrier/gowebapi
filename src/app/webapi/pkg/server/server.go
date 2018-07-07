package server

import (
	"fmt"
	"net/http"
)

// Config stores the hostname and port number
type Config struct {
	Hostname           string `json:"Hostname"`           // Server name.
	UseHTTP            bool   `json:"UseHTTP"`            // Listen on HTTP.
	UseHTTPS           bool   `json:"UseHTTPS"`           // Listen on HTTPS.
	ForceHTTPSRedirect bool   `json:"ForceHTTPSRedirect"` // Redirect HTTP to HTTPS.
	HTTPPort           int    `json:"HTTPPort"`           // HTTP port.
	HTTPSPort          int    `json:"HTTPSPort"`          // HTTPS port.
	CertFile           string `json:"CertFile"`           // HTTPS certificate.
	KeyFile            string `json:"KeyFile"`            // HTTPS private key.
}

// ILog provides logging capabilities.
type ILog interface {
	Fatalf(format string, v ...interface{})
	Printf(format string, v ...interface{})
}

// IServer provides HTTP server capabilities.
type IServer interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
}

// Run starts the HTTP and/or HTTPS listener
func (c *Config) Run(httpServer IServer, httpsServer IServer, logger ILog) {
	if c.UseHTTP && c.UseHTTPS {
		go func() {
			c.StartHTTPS(httpsServer, logger)
		}()

		c.StartHTTP(httpServer, logger)
	} else if c.UseHTTP {
		c.StartHTTP(httpServer, logger)
	} else if c.UseHTTPS {
		c.StartHTTPS(httpsServer, logger)
	} else {
		logger.Printf("Config file does not specify a listener to start")
	}
}

// StartHTTP starts the HTTP listener.
func (c *Config) StartHTTP(s IServer, logger ILog) {
	addr := c.HTTPAddress()
	logger.Printf("%v", "Running HTTP "+addr)

	// Start the HTTP listener
	logger.Fatalf("%v", s.ListenAndServe())
}

// StartHTTPS starts the HTTPS listener.
func (c *Config) StartHTTPS(s IServer, logger ILog) {
	addr := c.HTTPSAddress()
	logger.Printf("%v", "Running HTTPS "+addr)

	// Start the HTTPS listener
	logger.Fatalf("%v", s.ListenAndServeTLS(c.CertFile, c.KeyFile))
}

// HTTPAddress returns the HTTP address.
func (c *Config) HTTPAddress() string {
	port := 80
	if c.HTTPPort != 0 {
		port = c.HTTPPort
	}
	return c.Hostname + ":" + fmt.Sprintf("%d", port)
}

// HTTPSAddress returns the HTTPS address.
func (c *Config) HTTPSAddress() string {
	port := 443
	if c.HTTPSPort != 0 {
		port = c.HTTPSPort
	}
	return c.Hostname + ":" + fmt.Sprintf("%d", port)
}

// RedirectToHTTPS will redirect HTTP to HTTPS.
func RedirectToHTTPS() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
	})
}
