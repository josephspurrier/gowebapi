package user

import (
	"app/webapi/pkg/structcopy"
	"net/http"

	"app/webapi/model"
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
	results := make(store.UserGroup, 0)
	_, err := u.FindAll(&results)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Copy to the response.
	arr := make([]model.UserIndexResponseData, 0)
	for _, u := range results {
		item := new(model.UserIndexResponseData)
		err = structcopy.ByTag(&u, "db", item, "json")
		if err != nil {
			return http.StatusInternalServerError, err
		}
		arr = append(arr, *item)
	}

	resp := new(model.UserIndexResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data = arr
	return p.Response.JSON(w, resp.Body)
}
