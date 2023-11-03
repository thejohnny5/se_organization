package services

import (
	"fmt"
)

// Returns claims moving forward (user_id)
func ValidateToken(token string) (string, error) {
	// Look up token in database
	// for now if token is equal to "user1"
	if token == "user1" {
		return "authuser1", nil
	}

	return "", fmt.Errorf("user is not logged in")
}
