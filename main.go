package main

import (
	"log"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/api" // Adjust the import path according to your module name
)

func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
