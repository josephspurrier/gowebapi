package root

import (
	"net/http"
)

// Index .
// swagger:route GET /v1 root RootIndex
//
// Display a hello message.
//
// Responses:
//   200: OKResponse
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	return p.Response.OK(w, "hello")
}
