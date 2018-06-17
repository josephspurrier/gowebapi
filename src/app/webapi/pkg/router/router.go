package router

import (
	"net/http"

	"github.com/matryer/way"
)

// Mux contains the router.
type Mux struct {
	router *way.Router
}

// New returns an instance of the router.
func New() *Mux {
	return &Mux{
		router: way.NewRouter(),
	}
}

// Instance returns the router.
func (m *Mux) Instance() *way.Router {
	return m.router
}

// ServeHTTP routes the incoming http.Request based on method and path
// extracting path parameters as it goes.
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

// Params returns a URL parameter.
func Params(r *http.Request, param string) string {
	return way.Param(r.Context(), param)
}
