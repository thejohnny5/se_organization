package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/services"
)

type Claims struct {
	UserID uint
}

type contextKey string

const claimsContextKey contextKey = "claims"

func (DB *DBClient) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		log.Printf("TokenString: %s", tokenString)
		// if no string
		// user not logged in error and redirect

		// validate token
		user_id, err := services.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid so add to context
		ctx := context.WithValue(r.Context(), claimsContextKey, Claims{UserID: user_id})

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
