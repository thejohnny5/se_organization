package main

import (
	"log"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/api" // Adjust the import path according to your module name
)

func main() {
	db, err := api.ConnectClient()
	if err != nil {
		panic("Can't connect to db")
	}
	router := api.NewRouter(db)
	log.Printf("Starting database on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
