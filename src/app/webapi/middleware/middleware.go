package middleware

import (
	"net/http"

	"app/webapi/middleware/cors"
	"app/webapi/middleware/logrequest"
)

// *****************************************************************************
// Middleware
// *****************************************************************************

// Wrap will return the http.Handler wrapped in middleware.
func Wrap(h http.Handler, logger logrequest.ILog,
	clock logrequest.IClock) http.Handler {
	// Log every request.
	lr := logrequest.New(logger, clock)
	h = lr.Handler(h)

	// CORS for the endpoints.
	h = cors.Handler(h)

	return h
}
