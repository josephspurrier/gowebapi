package auth

import (
	"net/http"
	"time"

	"app/webapi/model"
)

// Show will: Show an access token.
// swagger:route GET /v1/auth auth AuthShow
//
// Show an access token.
//
// Responses:
//   200: AuthShowResponse
func (p *Endpoint) Show(w http.ResponseWriter, r *http.Request) (int, error) {
	t, err := p.Token.Generate("1", 8*time.Hour)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp := new(model.AuthShowResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data.Token = t
	return p.Response.JSON(w, resp.Body)
}
