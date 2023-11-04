package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/models"
)

func (db *DBClient) CreateUser(w http.ResponseWriter, r *http.Request) {
	// decode body
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close() // Make sure to close body at end of call stack
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// unmarshall
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("user details: %+v", user)
	// parse input for empty string/fields
	if user.UserEmail == "" || user.UserID == "" {
		http.Error(w, errors.New("user_email and user_id are required").Error(), http.StatusBadRequest)
		return
	}
	// if empty return with error
	// check if user_id is already in database -> error

	// At this point we would need to verify JWT token from google but yeah.

	// insert entry into database
	result := db.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// set Bearer -> user1 for testing

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (db *DBClient) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Read jwt token from request by google
	// verify jwt token against private key
	// if not correct, fail
	// else set token on user browser
}
