package user

import (
	"net/http"

	"app/webapi/store"
)

// Index .
// swagger:route GET /v1/user user UserIndex
//
// Return all users.
//
// Security:
//   token:
//
// Responses:
//   200: UserIndexResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	// Create the store.
	u := store.NewUser(p.DB, p.Q)

	// Get all items.
	results := make(store.TUserGroup, 0)
	_, err := u.FindAll(&results)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Response returns 200.
	// swagger:response UserIndexResponse
	type response struct {
		// in: body
		Body struct {
			// Required: true
			Status string `json:"status"`
			// Required: true
			Data store.TUserGroup `json:"data"`
		}
	}

	resp := new(response)
	return p.Response.Results(w, &resp.Body, results)
}
