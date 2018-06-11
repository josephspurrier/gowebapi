package router

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"

	"github.com/justinas/alice"
)

var (
	r RouterInfo
)

const (
	params = "params"
)

// RouteInfo is the details
type RouterInfo struct {
	Router *httprouter.Router
}

// Set up the router
func init() {
	r.Router = httprouter.New()
}

// ReadConfig returns the information
func ReadConfig() RouterInfo {
	return r
}

// Instance returns the router
func Instance() *httprouter.Router {
	return r.Router
}

// Context returns the URL parameters
func Params(r *http.Request) httprouter.Params {
	return context.Get(r, params).(httprouter.Params)
}

// Chain returns handle with chaining using Alice
func Chain(fn http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	return Handler(alice.New(c...).ThenFunc(fn))
}
