package static

import (
	"net/http"
	"strings"
)

// Static displays static files.
func (p *Endpoint) Static(w http.ResponseWriter, r *http.Request) (int, error) {
	// Disable listing directories.
	if strings.HasSuffix(r.URL.Path, "/") {
		return http.StatusNotFound, nil
	}

	http.ServeFile(w, r, "static/"+r.URL.Path[1:])
	return http.StatusOK, nil
}
