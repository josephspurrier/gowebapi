package model

// AuthShowResponse returns 200.
// swagger:response AuthShowResponse
type AuthShowResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data struct {
			// Required: true
			Token string `json:"token"`
		} `json:"data"`
	}
}
