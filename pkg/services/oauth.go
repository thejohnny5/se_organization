package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var googleAuthConfig oauth2.Config = oauth2.Config{
	RedirectURL:  "http://127.0.0.1:8080/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
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

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

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

	log.Println(data)
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
