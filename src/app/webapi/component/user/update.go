package user

import (
	"errors"
	"net/http"

	"app/webapi/store"
)

// Update .
// swagger:route PUT /v1/user/{user_id} user UserUpdate
//
// Make changes to a user.
//
// Security:
//   token:
//
// Responses:
//   200: OKResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Update(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters UserUpdate
	type request struct {
		// in: path
		// x-example: USERID
		UserID string `json:"user_id" validate:"required"`
		// in: formData
		// Required: true
		FirstName string `json:"first_name" validate:"required"`
		// in: formData
		// Required: true
		LastName string `json:"last_name" validate:"required"`
		// in: formData
		// Required: true
		Email string `json:"email" validate:"required"`
		// in: formData
		// Required: true
		Password string `json:"password" validate:"required"`
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

	// Determine if the item exists.
	exists, err := u.ExistsByID(u, req.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("user not found")
	}

	// Encrypt the password.
	password, err := p.Password.HashString(req.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Update the item.
	err = u.Update(u.ID, req.FirstName, req.LastName, req.Email, password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.OK(w, "user updated")
}
