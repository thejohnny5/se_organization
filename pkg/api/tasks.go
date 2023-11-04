package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/models"
)

func (db *DBClient) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Grab claims value from context
	claims, err := GetClaims(r)
	if err != nil {
		log.Printf("error returning claims: %s", err)
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	log.Printf("claims: %+v", claims)
	// Get Tasks associated with user id i in
	var tasks []models.Task
	result := db.DB.Where("user_id = ?", claims.UserID).Find(&tasks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (db *DBClient) CreateTask(w http.ResponseWriter, r *http.Request) {
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
	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("task details: %+v", task)

	// set user_id to claims.UserID (or could check if they match before deciding)
	task.UserID = claims.UserID
	result := db.DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (db *DBClient) UpdateTask(w http.ResponseWriter, r *http.Request) {
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
	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("task details: %+v", task)

	// set user_id to claims.UserID (or could check if they match before deciding)
	task.UserID = claims.UserID
	result := db.DB.Save(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (db *DBClient) DeleteTask(w http.ResponseWriter, r *http.Request) {
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
	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("task details: %+v", task)

	// set user_id to claims.UserID (or could check if they match before deciding)
	recordToDelete := models.Task{ID: task.ID, UserID: claims.UserID}

	result := db.DB.Delete(&recordToDelete)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
