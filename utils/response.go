package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type (
	// object represents an object
	object map[string]interface{}

	// response represents a response
	response struct {
		code int
		Data interface{} `json:"data,omitempty"`
	}
)

// ServeJSON serves the response to writer as JSON
func (resp *response) ServeJSON(w http.ResponseWriter) {
	if resp.code == 0 {
		panic(errors.New("response status not defined"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.code)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

// ParseBody parses request body to given data struct
func ParseBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
