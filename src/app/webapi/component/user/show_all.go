package user

import (
	"net/http"

	"app/webapi/model"
	"app/webapi/pkg/structcopy"
	"app/webapi/store"
)

// ShowAll will: Show all users.
// swagger:route GET /v1/user user UserShowAll
//
// Show all users.
//
// Security:
//   token:
//
// Responses:
//   200: UserShowAllResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) ShowAll(w http.ResponseWriter, r *http.Request) (int, error) {
	// Create the DB store.
	u := store.NewUser(p.DB, p.Q)

	// Get all items.
	results := make(store.UserGroup, 0)
	_, err := u.FindAll(&results)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Copy the items to the JSON model.
	arr := make([]model.UserShowAllResponseData, 0)
	for _, u := range results {
		item := new(model.UserShowAllResponseData)
		err = structcopy.ByTag(&u, "db", item, "json")
		if err != nil {
			return http.StatusInternalServerError, err
		}
		arr = append(arr, *item)
	}

	// Send the response.
	resp := new(model.UserShowAllResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data = arr
	return p.Response.JSON(w, resp.Body)
}
