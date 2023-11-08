package api

import (
	"encoding/json"
	"net/http"

	"github.com/thejohnny5/se_organization/pkg/models"
)

type DropDownRequest struct {
	ID        uint   `json:"id"`
	TableType string `json:"table_type"`
	Text      string `json:"text"`
}
type DropDownDBHandler struct {
	DB *models.DBClient
}

func CreateDropdownDB(db *models.DBClient) *DropDownDBHandler {
	return &DropDownDBHandler{DB: db}
}

// TODO -> allow user to create his/her own dropdowns
// which would require changing this function to check
// for user_id from GetClaims

func (db *DropDownDBHandler) GetDropdownOptions(w http.ResponseWriter, r *http.Request) {
	var dd []models.Dropdown

	tableType := r.URL.Query().Get("tabletype")

	result := db.DB.DB.Where("user_id is NULL").Where("table_type=?", tableType).Find(&dd)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	dropdown_return := make([]DropDownRequest, len(dd))
	for i, drop := range dd {
		dropdown_return[i] = DropDownRequest{
			ID:        drop.ID,
			TableType: drop.TableType,
			Text:      drop.Text,
		}
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(dropdown_return)
}
