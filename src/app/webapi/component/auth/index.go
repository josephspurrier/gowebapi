package auth

import (
	"net/http"
	"time"

	"app/webapi/model"
)

// Index .
// swagger:route GET /v1/auth auth AuthIndex
//
// Get an access token.
//
// Responses:
//   200: AuthIndexResponse
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	t, err := p.Token.Generate("1", 8*time.Hour)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp := new(model.AuthIndexResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data.Token = t
	return p.Response.JSON(w, resp.Body)
}
