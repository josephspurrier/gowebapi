package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"app/webapi/internal/response"
	"app/webapi/internal/webtoken"
)

// Config contains the dependencies for the handler.
type Config struct {
	clock  IClock
	secret []byte
}

// IClock provides clock capabilities.
type IClock interface {
	Now() time.Time
}

// New returns a new loq request middleware.
func New(secret []byte, clock IClock) *Config {
	return &Config{
		clock:  clock,
		secret: secret,
	}
}

// Handler will require a JWT.
func (c *Config) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// JWT whitelist.
		whitelist := []string{
			"GET /v1",
			"GET /v1/auth",
		}

		// Determine if the page is in the JWT whitelist.
		if !isWhitelisted(r.Method, r.URL.Path, whitelist) {
			// Require JWT on all routes.
			bearer := r.Header.Get("Authorization")

			// If the token is missing, show an error.
			if len(bearer) < 8 || !strings.HasPrefix(bearer, "Bearer ") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				r := new(response.UnauthorizedResponse)
				r.Body.Status = http.StatusText(http.StatusUnauthorized)
				r.Body.Message = "authorization token is missing"
				err := json.NewEncoder(w).Encode(r.Body)
				if err != nil {
					w.Write([]byte(`{"status":"Internal Server Error","message":"problem encoding JSON"}`))
					return
				}
				return
			}

			token := webtoken.New(c.secret, c.clock)
			_, err := token.Verify(bearer[7:])
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				r := new(response.UnauthorizedResponse)
				r.Body.Status = http.StatusText(http.StatusUnauthorized)
				r.Body.Message = "authorization token is invalid"
				err := json.NewEncoder(w).Encode(r.Body)
				if err != nil {
					w.Write([]byte(`{"status":"Internal Server Error","message":"problem encoding JSON"}`))
					return
				}
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// isWhitelisted returns true if the request is in the whitelist. If an
// asterisk is found in the whitelist, allow all routes.
func isWhitelisted(method string, path string, arr []string) (found bool) {
	s := fmt.Sprintf("%v %v", method, path)
	for _, i := range arr {
		if i == "*" || s == i {
			return true
		}
	}
	return
}
