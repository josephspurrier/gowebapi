package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// Output is the response object.
type Output struct{}

// New returns a new response object.
func New() *Output {
	return &Output{}
}

// OK will write an OK status to the writer.
func (o *Output) OK(w http.ResponseWriter, message string) (int, error) {
	r := new(OKResponse)
	r.Body.Status = http.StatusText(http.StatusOK)
	r.Body.Message = message

	// Write the content.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(r.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Created will output a creation response to the writer.
func (o *Output) Created(w http.ResponseWriter, recordID string) (int, error) {
	r := new(CreatedResponse)
	r.Body.Status = http.StatusText(http.StatusCreated)
	r.Body.RecordID = recordID

	// Write the content.
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(r.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

// Results will output the data to the writer.
func (o *Output) Results(w http.ResponseWriter, body interface{}, data interface{}) (int, error) {
	if data != nil {
		// Ensure a pointer is passed in.
		v := reflect.ValueOf(body)
		if v.Kind() != reflect.Ptr {
			return http.StatusInternalServerError, fmt.Errorf("body types do not match - expected 'struct pointer' but got '%v'", v.Kind())
		}

		// Ensure a struct is passed in.
		v = reflect.Indirect(reflect.ValueOf(body))
		if v.Kind() != reflect.Struct {
			return http.StatusInternalServerError, fmt.Errorf("body types do not match - expected 'struct pointer' but got '%v pointer'", v.Kind())
		}

		// Loop through each field.
		keys := v.Type()
		for j := 0; j < v.NumField(); j++ {
			field := v.Field(j)
			tag := keys.Field(j).Tag

			// Set the "status" field.
			if tag.Get("json") == "status" {
				if field.Kind() == reflect.String {
					v.Field(j).SetString(http.StatusText(http.StatusOK))
				} else {
					return http.StatusInternalServerError, fmt.Errorf("data types do not match for 'status' - expected '%v' but got '%v'", reflect.String, field.Type())
				}
			}

			// Set the "data" field.
			if tag.Get("json") == "data" {
				dataType := reflect.TypeOf(data)
				if field.Type() == dataType {
					v.Field(j).Set(reflect.ValueOf(data))
				} else {
					return http.StatusInternalServerError, fmt.Errorf("data types do not match for 'data' - expected %v but got %v", field.Type(), dataType)
				}

			}
		}
	}

	// Write the content.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// OKResponse returns 200.
// swagger:response OKResponse
type OKResponse struct {
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

// BadRequestResponse returns 400.
// swagger:response BadRequestResponse
type BadRequestResponse struct {
	OKResponse
}

// UnauthorizedResponse returns 401.
// swagger:response UnauthorizedResponse
type UnauthorizedResponse struct {
	OKResponse
}

// NotFoundResponse returns 404.
// swagger:response NotFoundResponse
type NotFoundResponse struct {
	OKResponse
}

// InternalServerErrorResponse returns 500.
// swagger:response InternalServerErrorResponse
type InternalServerErrorResponse struct {
	OKResponse
}
