package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/thejohnny5/se_organization/pkg/models"
	"github.com/thejohnny5/se_organization/pkg/services"
)

func NewRouter(DB *models.DBClient) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Will be for static assets eventually
	router.PathPrefix("/assets/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/assets/"):]
		fullPath := "/app/frontend/assets/" + path
		if strings.HasSuffix(path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		http.ServeFile(w, r, fullPath)
	})

	// userHandler := CreateUserDB(DB)
	authHandler := CreateAuthDB(DB)
	dropdownHandler := CreateDropdownDB(DB)
	jobHandler := CreateJobsDB(DB)
	documentHandler := CreateFileDB(DB)
	oauthHandler := services.CreateAuthDB(DB)
	// router.HandleFunc("/api/user/create", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/oauth/google/login", services.HandleGoogleLogin)
	router.HandleFunc("/oauth/google/callback", oauthHandler.HandleGoogleCallback)
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Setup API router to always first verify JWT Token
	apiRouter.Use(authHandler.AuthenticationMiddleware)
	apiRouter.HandleFunc("/validate", validate).Methods("GET")
	apiRouter.HandleFunc("/logout", logout).Methods("POST")
	// API Routes
	// apiRouter.HandleFunc("/tasks", DB.GetTasks).Methods("GET")
	// apiRouter.HandleFunc("/tasks", DB.CreateTask).Methods("POST")
	// apiRouter.HandleFunc("/tasks", DB.UpdateTask).Methods("PUT")
	// apiRouter.HandleFunc("/tasks/{id}", DB.DeleteTask).Methods("DELETE")
	// Dropdowns
	apiRouter.HandleFunc("/dropdown", dropdownHandler.GetDropdownOptions).Methods("GET")

	apiRouter.HandleFunc("/job", jobHandler.GetJobs).Methods("GET")
	apiRouter.HandleFunc("/job", jobHandler.CreateJob).Methods("POST")
	apiRouter.HandleFunc("/job", jobHandler.UpdateJob).Methods("PUT")
	apiRouter.HandleFunc("/job/{id}", jobHandler.DeleteJob).Methods("DELETE")
	apiRouter.HandleFunc("/job/csvdownload", jobHandler.DownloadCSV).Methods("GET")
	apiRouter.HandleFunc("/job/csvupload", jobHandler.UploadCSV).Methods("POST")
	apiRouter.HandleFunc("/document", documentHandler.UploadFile).Methods("POST")
	apiRouter.HandleFunc("/document", documentHandler.GetDocuments).Methods("GET")
	apiRouter.HandleFunc("/document/{id}/download", documentHandler.DownloadDoc).Methods("GET")

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve index.html for all non-API routes
		http.ServeFile(w, r, "/app/frontend/index.html")
	})
	return router
}

func validate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}
