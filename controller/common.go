package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Status indicates the status of the request response
type Status string

const (
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
	StatusUnknown Status = "unknown"
)

// Response defines format for the response of a http request
type Response struct {
	HTTPSTatus int         `json:"-"`
	Status     Status      `json:"status"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// WriteJson writes a json response to the http response
func (r Response) WriteJson(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.HTTPSTatus)

	return json.NewEncoder(w).Encode(r)
}

// handleNotFound handles all not found requests
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	Response{
		HTTPSTatus: http.StatusNotFound,
		Status:     StatusUnknown,
		Message: fmt.Sprintf(
			"Unknown Resource: %s -> %s",
			r.Method,
			r.URL.RequestURI(),
		),
	}.WriteJson(w)
}

// handleMethodNotAllowed handles all not allowed requests
func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	Response{
		HTTPSTatus: http.StatusMethodNotAllowed,
		Status:     StatusUnknown,
		Message: fmt.Sprintf(
			"Unknown Action: %s -> %s",
			r.Method,
			r.URL.RequestURI(),
		),
	}.WriteJson(w)
}
