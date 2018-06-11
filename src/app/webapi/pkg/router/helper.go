package router

import (
	"net/http"
)

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func Delete(path string, fn http.HandlerFunc) {
	r.Router.Handle("DELETE", path, fn)
}

// Get is a shortcut for router.Handle("GET", path, handle)
func Get(path string, fn http.HandlerFunc) {
	r.Router.Handle("GET", path, fn)
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func Head(path string, fn http.HandlerFunc) {
	r.Router.Handle("HEAD", path, fn)
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func Options(path string, fn http.HandlerFunc) {
	r.Router.Handle("OPTIONS", path, fn)
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func Patch(path string, fn http.HandlerFunc) {
	r.Router.Handle("PATCH", path, fn)
}

// Post is a shortcut for router.Handle("POST", path, handle)
func Post(path string, fn http.HandlerFunc) {
	r.Router.Handle("POST", path, fn)
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func Put(path string, fn http.HandlerFunc) {
	r.Router.Handle("PUT", path, fn)
}
