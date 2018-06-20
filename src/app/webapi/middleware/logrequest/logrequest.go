package logrequest

import (
	"log"
	"net/http"
	"os"
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

// New returns a new logrequest config.
func New() *Config {
	return &Config{}
}

// SetClock will set the clock.
func (c *Config) SetClock(clock IClock) {
	c.clock = clock
}

// SetLog will set the logger.
func (c *Config) SetLog(l ILog) {
	c.log = l
}

// Handler will log the HTTP requests.
func (c *Config) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the current time.
		now := time.Now()
		if c.clock != nil {
			now = c.clock.Now()
		}

		// Set the logger.
		var l ILog
		l = log.New(os.Stderr, "", log.LstdFlags)
		if c.log != nil {
			l = c.log
		}

		l.Printf("%v %v %v %v", now.Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
