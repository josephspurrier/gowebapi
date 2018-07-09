package middleware

import (
	"net/http"

	"app/webapi/middleware/cors"
	"app/webapi/middleware/jwt"
	"app/webapi/middleware/logrequest"
)

// *****************************************************************************
// Middleware
// *****************************************************************************

// Wrap will return the http.Handler wrapped in middleware.
func Wrap(h http.Handler, l logrequest.ILog, secret []byte) http.Handler {
	// JWT whitelist.
	whitelist := []string{
		"GET /v1",
		"GET /v1/auth",
	}

	// JWT validation.
	token := jwt.New(secret, whitelist)
	h = token.Handler(h)

	// CORS for the endpoints.
	h = cors.Handler(h)

	// Log every request.
	lr := logrequest.New()
	lr.SetLog(l)
	h = lr.Handler(h)

	return h
}
