package model

// GenericResponse returns any status code.
type GenericResponse struct {
	// in: body
	Body struct {
		// Status contains the string of the HTTP status.
		//
		// Required: true
		Status string `json:"status"`
		// Message can contain a user friendly message.
		Message string `json:"message,omitempty"`
	}
}

// CreatedResponse returns 201.
// swagger:response CreatedResponse
type CreatedResponse struct {
	// in: body
	Body struct {
		// Status contains the string of the HTTP status.
		//
		// Required: true
		Status string `json:"status"`
		// RecordID can be used for returning the ID from a row.
		RecordID string `json:"record_id,omitempty"`
	}
}

// OKResponse returns 200.
// swagger:response OKResponse
type OKResponse struct {
	GenericResponse
}

// BadRequestResponse returns 400.
// swagger:response BadRequestResponse
type BadRequestResponse struct {
	GenericResponse
}

// UnauthorizedResponse returns 401.
// swagger:response UnauthorizedResponse
type UnauthorizedResponse struct {
	GenericResponse
}

// InternalServerErrorResponse returns 500.
// swagger:response InternalServerErrorResponse
type InternalServerErrorResponse struct {
	GenericResponse
}
