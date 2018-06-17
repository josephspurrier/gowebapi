package logrequest

import (
	"net/http"
	"time"
)

// Config contains the dependencies for the handler.
type Config struct {
	log   ILog
	clock IClock
}

// ILog provides logging capabilities.
type ILog interface {
	Printf(format string, v ...interface{})
}

// IClock provides clock capabilities.
type IClock interface {
	Now() time.Time
}

// New returns a new loq request middleware.
func New(l ILog, clock IClock) *Config {
	return &Config{
		log:   l,
		clock: clock,
	}
}

// Handler will log the HTTP requests.
func (c *Config) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.log.Printf("%v %v %v %v", c.clock.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
