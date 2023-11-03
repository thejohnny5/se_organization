package api

import (
	"encoding/json"
	"net/http"
)

func (db *DBClient) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []string{"Task1", "Task2", "Task3"}
	json.NewEncoder(w).Encode(tasks)
}
