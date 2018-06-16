package component

import (
	"encoding/json"
	"log"
	"net/http"

	"app/webapi/internal/response"
)

// H is used to wrapper all endpoint functions so they work with generic
// routers.
type H func(http.ResponseWriter, *http.Request) (int, error)

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn(w, r)
	// Handle only errors.
	if status >= 400 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		r := new(response.OKResponse)
		r.Body.Status = http.StatusText(status)
		if err != nil {
			r.Body.Message = err.Error()
		}

		err := json.NewEncoder(w).Encode(r.Body)
		if err != nil {
			w.Write([]byte(`{"status":"Internal Server Error","message":"problem encoding JSON"}`))
			return
		}
	}

	// Only output 500 errors.
	if status >= 500 {
		if err != nil {
			log.Println(err)
		}
	}
}
