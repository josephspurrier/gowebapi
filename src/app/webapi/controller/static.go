package controller

import (
	"net/http"
	"strings"

	"app/webapi/pkg/router"
)

func init() {
	// Serve static files, no directory browsing.
	router.Get("/static/*filepath", Static)
}

// Static displays static files.
func Static(w http.ResponseWriter, r *http.Request) {
	// Disable listing directories.
	if strings.HasSuffix(r.URL.Path, "/") {
		Error404(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
