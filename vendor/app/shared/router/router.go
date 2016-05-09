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
	return context.Get(r, "params").(httprouter.Params)
}

// Chain returns handle with chaining using Alice
func Chain(fn http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	return Handler(alice.New(c...).ThenFunc(fn))
}

// HandlerFunc accepts the name of a function so you don't have to wrap it with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.HandlerFunc(controller.Index))
// Source: http://nicolasmerouze.com/guide-routers-golang/
func HandlerFunc(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
}

// Handler accepts a handler to make it compatible with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.Handler(http.HandlerFunc(controller.Index)))
// Source: http://nicolasmerouze.com/guide-routers-golang/
func Handler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
}
