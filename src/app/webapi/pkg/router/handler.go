package router

import (
	"net/http"
)

// Handler is used to wrapper all endpoint functions so they work with generic
// routers.
type Handler func(http.ResponseWriter, *http.Request) (int, error)

// ServeHTTP is a settable function that receives the status and error from
// the function call.
var ServeHTTP = func(w http.ResponseWriter, r *http.Request, status int,
	err error) {
	if status >= 400 {
		if err != nil {
			http.Error(w, err.Error(), status)
		} else {
			http.Error(w, "", status)
		}
	}
}

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn(w, r)
	ServeHTTP(w, r, status, err)
}
