package response

// CreatedResponse returns 201.
// swagger:response CreatedResponse
type CreatedResponse struct {
	// in: body
	Body struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		// RecordID can be used for returning the ID from a row.
		RecordID int64 `json:"record_id,omitempty"`
	}
}

// ErrorResponse is a standard error response.
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		Status  string `json:"status"`
		Message string `json:"message,omitempty"`
	}
}

// BadRequestResponse returns 400.
// swagger:response BadRequestResponse
type BadRequestResponse struct {
	ErrorResponse
}

// UnauthorizedResponse returns 401.
// swagger:response UnauthorizedResponse
type UnauthorizedResponse struct {
	ErrorResponse
}

// NotFoundResponse returns 404.
// swagger:response NotFoundResponse
type NotFoundResponse struct {
	ErrorResponse
}

// InternalServerErrorResponse returns 500.
// swagger:response InternalServerErrorResponse
type InternalServerErrorResponse struct {
	ErrorResponse
}
