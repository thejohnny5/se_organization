package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Setup API router to always first verify JWT Token
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(AuthenticationMiddleware)
	// API Routes
	apiRouter.HandleFunc("/tasks").Methods("GET")
	apiRouter.HandleFunc("/job_applications").Methods("GET")
	return router
}
