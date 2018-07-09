package user

import (
	"errors"
	"net/http"

	"app/webapi/store"
)

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
	// Create the DB store.
	u := store.NewUser(p.DB, p.Q)

	// Delete all items.
	count, err := u.DeleteAll(u)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count < 1 {
		return http.StatusBadRequest, errors.New("no users to delete")
	}

	return p.Response.OK(w, "users deleted")
}
