package user

import (
	"errors"
	"net/http"

	"app/webapi/store"
)

// *****************************************************************************
// Read
// *****************************************************************************

// Show .
// swagger:route GET /v1/user/{user_id} user UserShow
//
// Return one user.
//
// Security:
//   token:
//
// Responses:
//   200: UserShowResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Show(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters UserShow
	type request struct {
		// in: path
		// x-example: USERID
		UserID string `json:"user_id" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(req, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(req); err != nil {
		return http.StatusBadRequest, err
	}

	// Create the store.
	u := store.NewUser(p.DB, p.Q)

	// Get a user.
	exists, err := u.FindOneByID(u, req.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("item not found")
	}

	// Response returns 200.
	// swagger:response UserShowResponse
	type response struct {
		// in: body
		Body struct {
			// Required: true
			Status string `json:"status"`
			// Required: true
			Data []store.User `json:"data"`
		}
	}

	resp := new(response)
	return p.Response.Results(w, &resp.Body, []store.User{*u})
}
