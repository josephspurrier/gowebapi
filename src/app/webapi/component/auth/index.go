package auth

import (
	"net/http"
	"time"
)

// Index .
// swagger:route GET /v1/auth auth AuthIndex
//
// Get an access token.
//
// Responses:
//   200: AuthIndexResponse
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	t, err := p.Token.Generate("jsmith", 8*time.Hour)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Response returns 200.
	// swagger:response AuthIndexResponse
	type response struct {
		// in: body
		Body struct {
			// Required: true
			Status string `json:"status"`
			// Required: true
			Data struct {
				// Required: true
				Token string `json:"token"`
			} `json:"data"`
		}
	}

	resp := new(response)
	resp.Body.Data.Token = t
	return p.Response.Results(w, &resp.Body, nil)
}
