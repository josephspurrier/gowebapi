package user

import (
	"errors"
	"net/http"

	"app/webapi/model"
	"app/webapi/pkg/structcopy"
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

	// Create the DB store.
	u := store.NewUser(p.DB, p.Q)

	// Get an item by ID.
	exists, err := u.FindOneByID(u, req.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("user not found")
	}

	// Copy the items to the JSON model.
	arr := make([]model.UserShowResponseData, 0)
	item := new(model.UserShowResponseData)
	err = structcopy.ByTag(u, "db", item, "json")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	arr = append(arr, *item)

	// Send the response.
	resp := new(model.UserShowResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data = arr
	return p.Response.JSON(w, resp.Body)
}
