package router

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (m *Mux) Delete(path string, fn Handler) {
	m.router.Handle("DELETE", path, fn)
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (m *Mux) Get(path string, fn Handler) {
	m.router.Handle("GET", path, fn)
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (m *Mux) Head(path string, fn Handler) {
	m.router.Handle("HEAD", path, fn)
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (m *Mux) Options(path string, fn Handler) {
	m.router.Handle("OPTIONS", path, fn)
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (m *Mux) Patch(path string, fn Handler) {
	m.router.Handle("PATCH", path, fn)
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (m *Mux) Post(path string, fn Handler) {
	m.router.Handle("POST", path, fn)
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (m *Mux) Put(path string, fn Handler) {
	m.router.Handle("PUT", path, fn)
}
