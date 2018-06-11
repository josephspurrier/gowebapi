package route

import (
	"net/http"

	"app/webapi/middleware/cors"
	"app/webapi/middleware/logrequest"
	"app/webapi/pkg/router"
)

// LoadHTTPS will load the HTTP routes and middleware.
func LoadHTTPS() http.Handler {
	//return middleware(routes())
	return middleware(router.Instance())
}

// LoadHTTP will load the HTTPS routes and middleware.
func LoadHTTP() http.Handler {
	//return middleware(routes())
	return middleware(router.Instance())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// Optional method to make it easy to redirect from HTTP to HTTPS.
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {
	// Log every request.
	h = logrequest.Handler(h)

	// Cors for swagger-ui.
	h = cors.Handler(h)

	return h
}
