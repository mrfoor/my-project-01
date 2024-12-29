package utils

import (
	"encoding/json" // Fixed typo in the import
	"io/ioutil"
	"net/http"
)

// ParseBody reads the request body and unmarshals it into the given interface
func ParseBody(r *http.Request, x interface{}) {
	// Read the body of the HTTP request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Log or handle the error as needed (optional)
		return
	}
	// Attempt to unmarshal the body into the provided interface
	if err := json.Unmarshal(body, x); err != nil {
		// Log or handle the error as needed (optional)
		return
	}
}
