package user

import (
	"errors"
	"net/http"

	"app/webapi/store"
)

// Create .
// swagger:route POST /v1/user user UserCreate
//
// Create a user.
//
// Security:
//   token:
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters UserCreate
	type request struct {
		// in: formData
		// Required: true
		FirstName string `json:"first_name" validate:"required"`
		// in: formData
		// Required: true
		LastName string `json:"last_name" validate:"required"`
		// in: formData
		// Required: true
		Email string `json:"email" validate:"required,email"`
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

	// Check for existing item.
	exists, _, err := u.ExistsByField(u, "email", req.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if exists {
		return http.StatusBadRequest, errors.New("user already exists")
	}

	// Encrypt the password.
	password, err := p.Password.HashString(req.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Create the item.
	ID, err := u.Create(req.FirstName, req.LastName, req.Email, password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.Created(w, ID)
}
