package model

// UserShowAllResponse returns 200.
// swagger:response UserShowAllResponse
type UserShowAllResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data []UserShowAllResponseData `json:"data"`
	}
}

// UserShowAllResponseData is the user data.
type UserShowAllResponseData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	StatusID  uint8  `json:"status_id"`
}
