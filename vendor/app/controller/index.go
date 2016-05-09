package controller

import (
	"net/http"

	"app/shared/response"
	"app/shared/router"
)

func init() {
	// Main page
	router.Instance().GET("/", router.HandlerFunc(Index))
}

func Index(w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "welcome", 0, nil)
}
