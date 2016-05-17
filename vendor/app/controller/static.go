package controller

import (
	"net/http"
	"strings"

	"app/shared/router"
)

func init() {
	// Required so the trailing slash is not redirected
	router.Instance().RedirectTrailingSlash = false

	// Serve static files, no directory browsing
	router.Get("/static/*filepath", Static)
}

// Static maps static files
func Static(w http.ResponseWriter, r *http.Request) {
	// Disable listing directories
	if strings.HasSuffix(r.URL.Path, "/") {
		Error404(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
