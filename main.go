package main

import (
	"log"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/api" // Adjust the import path according to your module name
	"github.com/thejohnny5/se_organization/pkg/models"
)

func main() {
	db, err := models.ConnectClient()
	if err != nil {
		panic("Can't connect to db")
	}
	router := api.NewRouter(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}
