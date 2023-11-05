package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thejohnny5/se_organization/pkg/models"
)

func (db *DBClient) GetJobs(w http.ResponseWriter, r *http.Request) {
	// Grab claims value from context
	claims, err := GetClaims(r)
	if err != nil {
		log.Printf("error returning claims: %s", err)
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	log.Printf("claims: %+v", claims)
	// Get jobs associated with user id i in
	var jobs []models.JobApplication
	result := db.DB.Where("user_id = ?", claims.UserID).Find(&jobs)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(jobs)
}

func (db *DBClient) CreateJob(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	// get request body decodeded
	body, berr := io.ReadAll(r.Body)
	defer r.Body.Close()
	if berr != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var job models.JobApplication
	if err := json.Unmarshal(body, &job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set user_id to claims.UserID (or could check if they match before deciding)
	job.UserID = claims.UserID
	log.Printf("job details: %+v", job)

	result := db.DB.Create(&job)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

func (db *DBClient) UpdateJob(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	// get request body decodeded
	body, berr := io.ReadAll(r.Body)
	defer r.Body.Close()
	if berr != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var job models.JobApplication
	if err := json.Unmarshal(body, &job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("job details: %+v", job)

	// set user_id to claims.UserID (or could check if they match before deciding)
	job.UserID = claims.UserID
	result := db.DB.Save(&job)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

func (db *DBClient) DeleteJob(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}

	// extract vars
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is missing in parameters", http.StatusBadRequest)
		return
	}

	// Parse the ID to a uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// set user_id to claims.UserID (or could check if they match before deciding)
	recordToDelete := models.JobApplication{ID: uint(id), UserID: claims.UserID}

	result := db.DB.Delete(&recordToDelete)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
