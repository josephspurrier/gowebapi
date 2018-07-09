package response

import (
	"encoding/json"
	"net/http"

	"app/webapi/model"
)

// Output is the response object.
type Output struct{}

// New returns a new response object.
func New() *Output {
	return &Output{}
}

// JSON will output JSON to the writer.
func (o *Output) JSON(w http.ResponseWriter, body interface{}) (int, error) {
	// Write the content.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// OK will write an OK status to the writer.
func (o *Output) OK(w http.ResponseWriter, message string) (int, error) {
	r := new(model.OKResponse)
	r.Body.Status = http.StatusText(http.StatusOK)
	r.Body.Message = message

	// Write the content.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(r.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Created will output a creation response to the writer.
func (o *Output) Created(w http.ResponseWriter, recordID string) (int, error) {
	r := new(model.CreatedResponse)
	r.Body.Status = http.StatusText(http.StatusCreated)
	r.Body.RecordID = recordID

	// Write the content.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(r.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
