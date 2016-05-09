package route

import (
	"net/http"

	"app/route/middleware/cors"
	"app/route/middleware/logrequest"
	"app/shared/router"

	"github.com/gorilla/context"
)

// Load the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	//return middleware(routes())
	return middleware(router.Instance())
}

// Load the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	//return middleware(routes())
	return middleware(router.Instance())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// Optional method to make it easy to redirect from HTTP to HTTPS
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {
	// Log every request
	h = logrequest.Handler(h)

	// Cors for swagger-ui
	h = cors.Handler(h)

	// Clear handler for Gorilla Context
	h = context.ClearHandler(h)

	return h
}
