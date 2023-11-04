package services

import (
	"fmt"
)

// Returns claims moving forward (user_id)
func ValidateToken(token string) (uint, error) {
	// Look up token in database
	// for now if token is equal to "user1"
	if token == "Bearer user1" {
		return 1, nil
	}

	return 0, fmt.Errorf("user is not logged in")
}
