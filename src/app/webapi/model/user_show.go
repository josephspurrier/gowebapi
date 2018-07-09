package model

// UserShowResponse returns 200.
// swagger:response UserShowResponse
type UserShowResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data []UserShowResponseData `json:"data"`
	}
}

// UserShowResponseData is the user data.
type UserShowResponseData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	StatusID  uint8  `json:"status_id"`
}
