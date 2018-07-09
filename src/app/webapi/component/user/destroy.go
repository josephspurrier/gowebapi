package user

import (
	"errors"
	"net/http"

	"app/webapi/store"
)

// Destroy .
// swagger:route DELETE /v1/user/{user_id} user UserDestroy
//
// Delete a user.
//
// Security:
//   token:
//
// Responses:
//   200: OKResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Destroy(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters UserDestroy
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

	// Create the DB store.
	u := store.NewUser(p.DB, p.Q)

	// Delete the item.
	count, err := u.DeleteOneByID(u, req.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count < 1 {
		return http.StatusBadRequest, errors.New("user does not exist")
	}

	return p.Response.OK(w, "user deleted")
}
