package router

import (
	"net/http"
)

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (r *Info) Delete(path string, fn http.Handler) {
	r.router.Handle("DELETE", path, fn)
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (r *Info) Get(path string, fn http.Handler) {
	r.router.Handle("GET", path, fn)
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (r *Info) Head(path string, fn http.Handler) {
	r.router.Handle("HEAD", path, fn)
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (r *Info) Options(path string, fn http.Handler) {
	r.router.Handle("OPTIONS", path, fn)
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (r *Info) Patch(path string, fn http.Handler) {
	r.router.Handle("PATCH", path, fn)
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (r *Info) Post(path string, fn http.Handler) {
	r.router.Handle("POST", path, fn)
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (r *Info) Put(path string, fn http.Handler) {
	r.router.Handle("PUT", path, fn)
}
