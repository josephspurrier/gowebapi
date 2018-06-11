package router

import (
	"net/http"
)

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func Delete(path string, fn http.HandlerFunc) {
	r.Router.DELETE(path, HandlerFunc(fn))
}

// Get is a shortcut for router.Handle("GET", path, handle)
func Get(path string, fn http.HandlerFunc) {
	r.Router.GET(path, HandlerFunc(fn))
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func Head(path string, fn http.HandlerFunc) {
	r.Router.HEAD(path, HandlerFunc(fn))
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func Options(path string, fn http.HandlerFunc) {
	r.Router.OPTIONS(path, HandlerFunc(fn))
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func Patch(path string, fn http.HandlerFunc) {
	r.Router.PATCH(path, HandlerFunc(fn))
}

// Post is a shortcut for router.Handle("POST", path, handle)
func Post(path string, fn http.HandlerFunc) {
	r.Router.POST(path, HandlerFunc(fn))
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func Put(path string, fn http.HandlerFunc) {
	r.Router.PUT(path, HandlerFunc(fn))
}
