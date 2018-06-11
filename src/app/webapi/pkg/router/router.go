package router

import (
	"net/http"

	"github.com/matryer/way"
)

var (
	r Info
)

const (
	params = "params"
)

// Info contains the router.
type Info struct {
	Router *way.Router
}

// Set up the router.
func init() {
	r.Router = way.NewRouter()
}

// ReadConfig returns the information.
func ReadConfig() Info {
	return r
}

// Instance returns the router.
func Instance() *way.Router {
	return r.Router
}

// Params returns a URL parameter.
func Params(r *http.Request, param string) string {
	return way.Param(r.Context(), param)
}
