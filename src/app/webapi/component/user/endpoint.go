package user

import (
	"log"
	"net/http"
)

const (
	itemCreated      = "item created"
	itemExists       = "item already exists"
	itemNotFound     = "item not found"
	itemFound        = "item found"
	itemsFound       = "items found"
	itemsFindEmpty   = "no items to find"
	itemUpdated      = "item updated"
	itemDeleted      = "item deleted"
	itemsDeleted     = "items deleted"
	itemsDeleteEmpty = "no items to delete"

	friendlyError = "an error occurred, please try again later"
)

// *****************************************************************************
// Create
// *****************************************************************************

// Create .
// swagger:route POST /v1/user user UserCreate
//
// Add a user to the database.
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Create(w http.ResponseWriter, r *http.Request) {
	// swagger:parameters UserCreate
	type request struct {
		// in: formData
		FirstName string `json:"first_name" validate:"required"`
		// in: formData
		LastName string `json:"last_name" validate:"required"`
		// in: formData
		Email string `json:"email" validate:"required"`
		// in: formData
		Password string `json:"password" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(&req, r); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	} else if err = p.Bind.Validate(req); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create an instance of the user.
	m, err := NewRecord()
	if err != nil {
		p.Log.Printf("UUID Error: %v", err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Copy over the values.
	m.FirstName = req.FirstName
	m.LastName = req.LastName
	m.Email = req.Email
	m.Password = req.Password

	// Create the user in the database.
	count, err := m.Create(p.DB)
	if err == ErrExists {
		p.Response.SendError(w, http.StatusBadRequest, itemExists)
		return
	} else if err != nil {
		p.Log.Printf("%v", err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	p.Response.Send(w, http.StatusCreated, itemCreated, count, nil)
}

// *****************************************************************************
// Read
// *****************************************************************************

// Show .
// swagger:route GET /v1/user/{user_id} user UserShow
//
// Return one user from the database.
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Show(w http.ResponseWriter, r *http.Request) {
	// swagger:parameters UserShow
	type request struct {
		// in: path
		UserID string `json:"user_id" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(&req, r); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	} else if err = p.Bind.Validate(req); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get an item.
	entity, err := Read(p.DB, req.UserID)
	if err == ErrNoResult {
		p.Response.Send(w, http.StatusOK, itemNotFound, 0, nil)
		return
	} else if err != nil {
		log.Println(err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	p.Response.Send(w, http.StatusOK, itemFound, 1, Group{*entity})
}

// Index .
// swagger:route GET /v1/user user UserIndex
//
// Returns all users in the database.
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) {
	// Get all items
	group, err := ReadAll(p.DB)
	if err != nil {
		log.Println(err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	} else if len(group) < 1 {
		p.Response.Send(w, http.StatusOK, itemsFindEmpty, len(group), nil)
		return
	}

	p.Response.Send(w, http.StatusOK, itemsFound, len(group), group)
}

// *****************************************************************************
// Update
// *****************************************************************************

// Update .
// swagger:route PUT /v1/user/{user_id} user UserUpdate
//
// Makes changes to a user in the database.
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Update(w http.ResponseWriter, r *http.Request) {
	// swagger:parameters UserUpdate
	type request struct {
		// in: path
		UserID string `json:"user_id" validate:"required"`
		// in: formData
		FirstName string `json:"first_name" validate:"required"`
		// in: formData
		LastName string `json:"last_name" validate:"required"`
		// in: formData
		Email string `json:"email" validate:"required"`
		// in: formData
		Password string `json:"password" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(&req, r); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	} else if err = p.Bind.Validate(req); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get an item.
	m, err := Read(p.DB, req.UserID)
	if err == ErrNoResult {
		p.Response.Send(w, http.StatusOK, itemNotFound, 0, nil)
		return
	} else if err != nil {
		log.Println(err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Copy over the values.
	m.FirstName = req.FirstName
	m.LastName = req.LastName
	m.Email = req.Email
	m.Password = req.Password

	// Update item
	count, err := m.Update(p.DB)
	if err == ErrNotExist {
		p.Response.SendError(w, http.StatusBadRequest, itemNotFound)
		return
	} else if err != nil {
		log.Println(err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	p.Response.Send(w, http.StatusCreated, itemUpdated, count, nil)
}

// *****************************************************************************
// Delete
// *****************************************************************************

// Destroy .
// swagger:route DELETE /v1/user/{user_id} user UserDestroy
//
// Delete a user from the database.
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Destroy(w http.ResponseWriter, r *http.Request) {
	// swagger:parameters UserDestroy
	type request struct {
		// in: path
		UserID string `json:"user_id" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(&req, r); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	} else if err = p.Bind.Validate(req); err != nil {
		p.Response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Delete an item.
	count, err := Delete(p.DB, req.UserID)
	if err != nil {
		log.Println(err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	} else if count < 1 {
		p.Response.Send(w, http.StatusOK, itemNotFound, count, nil)
		return
	}

	p.Response.Send(w, http.StatusOK, itemDeleted, count, nil)
}

// DestroyAll .
// swagger:route DELETE /v1/user user UserDestroyAll
//
// Deletes all users from the database.
//
// Responses:
//   200: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) DestroyAll(w http.ResponseWriter, r *http.Request) {
	// Delete all items
	count, err := DeleteAll(p.DB)
	if err != nil {
		log.Println(err)
		p.Response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	} else if count < 1 {
		p.Response.Send(w, http.StatusOK, itemsDeleteEmpty, count, nil)
		return
	}

	p.Response.Send(w, http.StatusOK, itemsDeleted, count, nil)
}
