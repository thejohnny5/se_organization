package api

import (
	"errors"
	"log"
	"net/http"
)

// Global Error
type APIError struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}

func (e *APIError) Error() string {
	return e.ErrorMessage
}

// Handle Error
func HandleError(w http.ResponseWriter, err error) {
	var apiErr *APIError
	// if error is API Error
	if errors.As(err, &apiErr) {
		log.Printf("Error: %s", apiErr.Error())
		http.Error(w, apiErr.Error(), apiErr.StatusCode)
	} else {
		log.Printf("Error: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
