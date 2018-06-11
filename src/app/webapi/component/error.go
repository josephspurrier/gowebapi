package component

import (
	"net/http"
)

// NotFoundHandler is the standard 404 handler.
type NotFoundHandler struct {
	Response IResponse
}

// Error404 - Page Not Found.
func (p *NotFoundHandler) Error404(w http.ResponseWriter, r *http.Request) {
	Error404(p.Response, w, r)
}

// Error404 - Page Not Found.
func Error404(resp IResponse, w http.ResponseWriter, r *http.Request) {
	resp.SendError(w, http.StatusNotFound, "not found")
}

// Error405 - Method Not Allowed.
func Error405(resp IResponse, w http.ResponseWriter, r *http.Request) {
	resp.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
}

// Error500 - Internal Server Error.
func Error500(resp IResponse, w http.ResponseWriter, r *http.Request) {
	resp.SendError(w, http.StatusInternalServerError, "internal server error")
}
