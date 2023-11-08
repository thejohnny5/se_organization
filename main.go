package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/thejohnny5/se_organization/pkg/api" // Adjust the import path according to your module name
	"github.com/thejohnny5/se_organization/pkg/models"
	"github.com/thejohnny5/se_organization/pkg/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("not using .env file")
	}
	services.GetGoogleAuthConfig()
	db, err := models.ConnectClient()
	if err != nil {
		panic("Can't connect to db")
	}
	router := api.NewRouter(db)
	log.Printf("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
