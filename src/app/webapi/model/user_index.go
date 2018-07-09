package model

// UserIndexResponse returns 200.
// swagger:response UserIndexResponse
type UserIndexResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data []UserIndexResponseData `json:"data"`
	}
}

// UserIndexResponseData is the user data.
type UserIndexResponseData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	StatusID  uint8  `json:"status_id"`
}
