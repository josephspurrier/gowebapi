package controller

import (
	"net/http"

	"app/webapi/pkg/response"
	"app/webapi/pkg/router"
)

func init() {
	// Main page.
	router.Get("/", Index)
}

// Index displays an ok message.
func Index(w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "ok", 0, nil)
}
