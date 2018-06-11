package static

import (
	"net/http"
	"strings"

	"app/webapi/component"
)

// Static displays static files.
func (p *Endpoint) Static(w http.ResponseWriter, r *http.Request) {
	// Disable listing directories.
	if strings.HasSuffix(r.URL.Path, "/") {
		component.Error404(p.Response, w, r)
		return
	}

	http.ServeFile(w, r, "static/"+r.URL.Path[1:])
}
