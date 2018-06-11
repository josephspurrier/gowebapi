package controller

import (
	"net/http"

	"app/webapi/pkg/response"
	"app/webapi/pkg/router"
)

func init() {
	// 404 Page.
	var e404 http.HandlerFunc = Error404
	router.Instance().NotFound = e404
}

// Error404 - Page Not Found.
func Error404(w http.ResponseWriter, r *http.Request) {
	response.SendError(w, http.StatusNotFound, "not found")
}

// Error405 - Method Not Allowed.
func Error405(w http.ResponseWriter, r *http.Request) {
	response.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
}

// Error500 - Internal Server Error.
func Error500(w http.ResponseWriter, r *http.Request) {
	response.SendError(w, http.StatusInternalServerError, "internal server error")
}
