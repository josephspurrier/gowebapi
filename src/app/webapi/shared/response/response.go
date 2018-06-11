package response

import (
	"encoding/json"
	"net/http"
)

// Core Response
type Core struct {
	Status  http.ConnState `json:"status"`
	Message string         `json:"message"`
}

// Change Response
type Change struct {
	Status   http.ConnState `json:"status"`
	Message  string         `json:"message"`
	Affected int            `json:"affected"`
}

// Retrieve Response
type Retrieve struct {
	Status  http.ConnState `json:"status"`
	Message string         `json:"message"`
	Count   int            `json:"count"`
	Results interface{}    `json:"results"`
}

// SendError calls Send by without a count or results
func SendError(w http.ResponseWriter, status http.ConnState, message string) {
	Send(w, status, message, 0, nil)
}

// Send writes struct to the writer using a format
func Send(w http.ResponseWriter, status http.ConnState, message string, count int, results interface{}) {

	var i interface{}

	// Determine the best format
	if count < 1 {
		i = &Core{
			Status:  status,
			Message: message,
		}
	} else if results == nil {
		i = &Change{
			Status:   status,
			Message:  message,
			Affected: count,
		}
	} else {
		i = &Retrieve{
			Status:  status,
			Message: message,
			Count:   count,
			Results: results,
		}
	}

	js, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	w.Write(js)
}

// SendJSON writes a struct to the writer
func SendJSON(w http.ResponseWriter, i interface{}) {
	js, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
