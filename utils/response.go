package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type (
	// object represents an object
	object map[string]interface{}

	// Response represents a response
	Response struct {
		Code         int         `json:"-"`
		Data         interface{} `json:"data,omitempty"`
		ErrorMessage string      `json:"error_message,omitempty"`
	}
)

// ServeJSON serves the response to writer as JSON
func (resp *Response) ServeJSON(w http.ResponseWriter) {
	if resp.Code == 0 {
		panic(errors.New("response status not defined"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

// ParseBody parses request body to given data struct
func ParseBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func HandleError(w http.ResponseWriter, statusCode int, message string) {
	resp := &Response{
		Code:         statusCode,
		ErrorMessage: message,
	}
	resp.ServeJSON(w)
}
