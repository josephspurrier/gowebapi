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
func Wrap(h http.Handler, logger logrequest.ILog, secret []byte) http.Handler {
	// JWT validation.
	token := jwt.New(secret)
	h = token.Handler(h)

	// CORS for the endpoints.
	h = cors.Handler(h)

	// Log every request.
	lr := logrequest.New()
	h = lr.Handler(h)

	return h
}
