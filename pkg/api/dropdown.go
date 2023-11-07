package api

import (
	"encoding/json"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/models"
)

// TODO -> allow user to create his/her own dropdowns
// which would require changing this function to check
// for user_id from GetClaims
func (db *DBClient) GetDropdownOptions(w http.ResponseWriter, r *http.Request) {
	var dd []models.Dropdown

	tableType := r.URL.Query().Get("tabletype")

	result := db.DB.Where("user_id is NULL").Where("table_type=?", tableType).Find(&dd)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dd)
}
