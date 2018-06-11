package router

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

var (
	r Info
)

const (
	params = "params"
)

// Info contains the router.
type Info struct {
	Router *httprouter.Router
}

// Set up the router.
func init() {
	r.Router = httprouter.New()
}

// ReadConfig returns the information.
func ReadConfig() Info {
	return r
}

// Instance returns the router.
func Instance() *httprouter.Router {
	return r.Router
}

// Params returns the URL parameters.
func Params(r *http.Request) httprouter.Params {
	return context.Get(r, params).(httprouter.Params)
}

// Chain returns handle with chaining using Alice.
func Chain(fn http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	return Handler(alice.New(c...).ThenFunc(fn))
}
