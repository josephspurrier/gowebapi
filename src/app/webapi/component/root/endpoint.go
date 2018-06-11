package root

import (
	"net/http"
)

// Index .
// swagger:route GET /v1 root RootIndex
//
// Displays an ok message.
//
// Responses:
//   200: RootIndexResponse
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) {
	// Response returns 200.
	// swagger:response RootIndexResponse
	type response struct {
		// in: body
		Body struct {
			Message string `json:"message"`
		}
	}

	p.Response.Send(w, http.StatusOK, "ok", 0, nil)
}
