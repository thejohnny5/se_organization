package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/thejohnny5/se_organization/pkg/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthDBHandler struct {
	DB *models.DBClient
}

func CreateAuthDB(db *models.DBClient) *OAuthDBHandler {
	return &OAuthDBHandler{DB: db}
}

type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var googleAuthConfig *oauth2.Config

func GetGoogleAuthConfig() *oauth2.Config {
	if googleAuthConfig == nil {
		googleAuthConfig = &oauth2.Config{
			RedirectURL:  "http://127.0.0.1:8080/oauth/google/callback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
	}
	return googleAuthConfig
}

func randomizeString(w http.ResponseWriter, numBytes uint) (string, error) {
	b := make([]byte, numBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.URLEncoding.EncodeToString(b)
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)
	return state, nil
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {

	oauthStateString, err := randomizeString(w, 16)
	if err != nil {
		http.Error(w, "Couldn't generate string", http.StatusInternalServerError)
		return
	}

	url := googleAuthConfig.AuthCodeURL(oauthStateString)
	log.Printf("google url: %s", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func (db *OAuthDBHandler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, err := r.Cookie("oauthstate")
	if err != nil {
		log.Println("no cookie present on user browser")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	data, err := getUserDataFromGoogle(r.FormValue("code"))

	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var user UserInfo
	if err := json.Unmarshal(data, &user); err != nil {
		log.Printf("Error handling unmarshaling: %v", err)
		http.Error(w, "couldn't unmarshalll data", http.StatusInternalServerError)
		return
	}
	log.Printf("user data: %+v", user)

	// handle the user data
	if err := db.handleUserLogin(w, &user); err != nil {
		log.Printf("Error handling user login: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// Check database and create user if needed
func (db *OAuthDBHandler) handleUserLogin(w http.ResponseWriter, user *UserInfo) error {
	var userModel models.User
	response := db.DB.DB.Where("email = ?", user.Email).First(&userModel)

	// Create a new user if not found in the DB
	if response.Error != nil {
		userModel.GoogleID = user.ID
		userModel.UserEmail = user.Email
		userModel.Picture = user.Picture
		userModel.VerifiedEmail = user.VerifiedEmail

		result := db.DB.DB.Create(&userModel)
		if result.Error != nil {
			return result.Error
		}
	}

	if err := db.startSession(w, &userModel); err != nil {
		return err
	}

	return nil
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := googleAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("failed getting user info %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %s", err.Error())
	}
	return contents, nil
}

func (db *OAuthDBHandler) startSession(w http.ResponseWriter, user *models.User) error {
	sessionToken := uuid.NewString() // or another secure random string
	sessionExpiration := time.Now().Add(12 * time.Hour)
	result := db.DB.DB.Exec("INSERT INTO sessions (session_token, user_id, expiration) VALUES (?, ?, ?)",
		sessionToken, user.ID, sessionExpiration)
	if result.Error != nil {
		return result.Error
	}
	// Set the session token as an HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  sessionExpiration,
		HttpOnly: true,
		Secure:   false, // change when using https
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}
