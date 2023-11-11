package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/thejohnny5/se_organization/pkg/models"
)

type AuthDBHandler struct {
	DB *models.DBClient
}

func CreateAuthDB(db *models.DBClient) *AuthDBHandler {
	return &AuthDBHandler{DB: db}
}

type Claims struct {
	UserID uint
}

type contextKey string

const claimsContextKey contextKey = "claims"

func (db *AuthDBHandler) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// tokenString, err := r.Cookie("session_token")
		// if err != nil {
		// 	//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }
		// // if no string
		// // user not logged in error and redirect
		// // validate token
		// user_id, err := db.ValidateToken(tokenString.Value)
		// if err != nil {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }

		// // Token is valid so add to context
		// ctx := context.WithValue(r.Context(), claimsContextKey, Claims{UserID: user_id})

		// ONLY FOR TESTING
		ctx := context.WithValue(r.Context(), claimsContextKey, Claims{UserID: 1})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaims(r *http.Request) (Claims, error) {
	claimsVal := r.Context().Value(claimsContextKey)
	if claimsVal == nil {
		return Claims{}, errors.New("no claims object")
	}
	// Type assertion
	claims, ok := claimsVal.(Claims)
	if !ok {
		return Claims{}, errors.New("Claims object is not correct type")
	}
	return claims, nil
}

// Returns claims moving forward (user_id)
func (db *AuthDBHandler) ValidateToken(token string) (uint, error) {
	// Look up token in database
	var session models.Session

	response := db.DB.DB.Where("session_token = ?", token).First(&session)

	if response.Error != nil {
		log.Printf("no user exists for session id: %s", token)

		return 0, errors.New("session does not exist")
	}

	// Check if the session has expired
	if session.Expiration.Before(time.Now()) {
		log.Printf("Session expired for token: %s", token)
		return 0, errors.New("unauthorized: session expired")
	}

	// Otherwise return user id
	return session.UserId, nil
}
