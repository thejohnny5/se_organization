package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thejohnny5/se_organization/pkg/models"
	"github.com/thejohnny5/se_organization/pkg/services"
)

func NewRouter(DB *models.DBClient) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// User creation
	router.Handle("/", http.FileServer(http.Dir("/app/frontend")))

	userHandler := CreateUserDB(DB)
	authHandler := CreateAuthDB(DB)
	dropdownHandler := CreateDropdownDB(DB)
	jobHandler := CreateJobsDB(DB)
	documentHandler := CreateFileDB(DB)
	oauthHandler := services.CreateAuthDB(DB)
	router.HandleFunc("/api/user/create", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/oauth/google/login", services.HandleGoogleLogin)
	router.HandleFunc("/oauth/google/callback", oauthHandler.HandleGoogleCallback)

	// Setup API router to always first verify JWT Token
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(authHandler.AuthenticationMiddleware)
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

	apiRouter.HandleFunc("/document", documentHandler.UploadFile).Methods("POST")
	apiRouter.HandleFunc("/document", documentHandler.GetDocuments).Methods("GET")
	apiRouter.HandleFunc("/document/{id}/download", documentHandler.DownloadDoc).Methods("GET")
	return router
}
