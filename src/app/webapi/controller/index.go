package controller

import (
	"net/http"

	"app/webapi/shared/response"
	"app/webapi/shared/router"
)

func init() {
	// Main page.
	router.Get("/", Index)
}

// Index displays an ok message.
func Index(w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "ok", 0, nil)
}
