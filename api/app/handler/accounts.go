package handler

import (
	"encoding/json"
	"net/http"
	"github.com/mikzuit/api/app/model"
	"github.com/jinzhu/gorm"
)

// Handler for Create accounts in POST method
func CreateAccount(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	// define and assing payload model
	record := model.Payload{}

	// decode payload
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&record); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&record.Data).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return 
	}
	respondJSON(w, http.StatusCreated, record)
}

func FetchAccount(db *gorm.DB, w http.ResponseWriter, r *http.Request){

}
