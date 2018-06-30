package user

import (
	"errors"
	"net/http"
)

// *****************************************************************************
// Create
// *****************************************************************************

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

	u := NewUser(p.DB, p.Q)

	// Check for existing user.
	exists, _, err := p.DB.RecordExistsString(func() (exists bool, ID string, err error) {
		return u.ExistsByField(u, "email", req.Email)
		//return ExistsEmail(p.DB, req.Email)
	})

	if err != nil {
		return http.StatusInternalServerError, err
	} else if exists {
		return http.StatusBadRequest, errors.New("user already exists")
	}

	// Create the user in the database.
	ID, err := p.DB.AddRecordString(func() (ID string, err error) {
		return u.Create(req.FirstName, req.LastName, req.Email, req.Password)
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.Created(w, ID)
}

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
		UserID string `json:"user_id" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(req, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(req); err != nil {
		return http.StatusBadRequest, err
	}

	// Get a user.
	u := NewUser(p.DB, p.Q)
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
			Data []TUser `json:"data"`
		}
	}

	resp := new(response)
	return p.Response.Results(w, &resp.Body, []TUser{*u})
}

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
	// Get all items.
	results := make(TUserGroup, 0)
	//FIXME: This code was commented out.
	//u := NewUser(p.DB, p.Q)

	/*_, err := p.DB.PaginatedResults(&results, func() (results interface{}, total int, err error) {
		return u.FindAll(&arr)
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}*/

	// Response returns 200.
	// swagger:response UserIndexResponse
	type response struct {
		// in: body
		Body struct {
			// Required: true
			Status string `json:"status"`
			// Required: true
			Data []TUser `json:"data"`
		}
	}

	resp := new(response)
	return p.Response.Results(w, &resp.Body, results)
}

// *****************************************************************************
// Update
// *****************************************************************************

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

	// Determine if the user exists.
	u := NewUser(p.DB, p.Q)
	exists, err := u.ExistsByID(u, req.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("user not found")
	}

	// Update item.
	err = u.Update(u.ID, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.OK(w, "user updated")
}

// *****************************************************************************
// Delete
// *****************************************************************************

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
		UserID string `json:"user_id" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(req, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(req); err != nil {
		return http.StatusBadRequest, err
	}

	// Delete an item.
	u := NewUser(p.DB, p.Q)
	count, err := u.DeleteOneByID(u, req.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count < 1 {
		return http.StatusBadRequest, errors.New("user does not exist")
	}

	return p.Response.OK(w, "user deleted")
}

// DestroyAll .
// swagger:route DELETE /v1/user user UserDestroyAll
//
// Delete all users.
//
// Security:
//   token:
//
// Responses:
//   200: OKResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) DestroyAll(w http.ResponseWriter, r *http.Request) (int, error) {
	// Delete all items.
	u := NewUser(p.DB, p.Q)
	count, err := u.DeleteAll(u)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count < 1 {
		return http.StatusBadRequest, errors.New("no users to delete")
	}

	return p.Response.OK(w, "users deleted")
}
