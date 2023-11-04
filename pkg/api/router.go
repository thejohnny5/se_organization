package api

import (
	"github.com/gorilla/mux"
)

func NewRouter(DB *DBClient) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// User creation
	router.HandleFunc("/api/user/create", DB.CreateUser).Methods("POST")
	// Setup API router to always first verify JWT Token
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(DB.AuthenticationMiddleware)
	// API Routes
	apiRouter.HandleFunc("/tasks", DB.GetTasks).Methods("GET")
	apiRouter.HandleFunc("/tasks", DB.CreateTask).Methods("POST")
	//apiRouter.HandleFunc("/job_applications").Methods("GET")
	return router
}
