package root

import (
	"net/http"
)

// Show wil: Display a hello message.
// swagger:route GET /v1 base BaseShow
//
// Display a hello message.
//
// Responses:
//   200: OKResponse
func (p *Endpoint) Show(w http.ResponseWriter, r *http.Request) (int, error) {
	return p.Response.OK(w, "hello")
}
