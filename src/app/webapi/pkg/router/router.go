package router

import (
	"net/http"

	"github.com/matryer/way"
)

// Info contains the router.
type Info struct {
	router *way.Router
}

// New returns an instance of the router.
func New() *Info {
	return &Info{
		router: way.NewRouter(),
	}
}

// Instance returns the router.
func (r *Info) Instance() *way.Router {
	return r.router
}

// Params returns a URL parameter.
func Params(r *http.Request, param string) string {
	return way.Param(r.Context(), param)
}
