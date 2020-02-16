package handler

import (
	"encoding/json"
	"net/http"
	"github.com/mikzuit/api/app/model"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
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
	respondJSON(w, http.StatusCreated, record.Data)
}

func FetchAccount(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	account := model.Account{}

	vars := mux.Vars(r)
	accountid := vars["accountid"]
		if err := db.First(&account, "aid = ?", accountid).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}

	respondJSON(w, http.StatusOK, account)
}

// Handler for DELETE accounts in POST method
func DeleteAccount(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	accountid := vars["accountid"]

	// get existent account or 404
	account := getAccountOr404(db, accountid, w, r)
	if account == nil {
		return
	}

	// Unscoped delete is a removed model, only delete just set the time deleted
	// this last one might cause unreliable testing
	if err := db.Unscoped().Delete(&account).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// Handler for LIST account with pagination
func GetPagination(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	id := vars["page_number"]
	limit := vars["page_size"]

	accounts := []model.Account{}

	db.Where("aid > ?", id).Order("ID ASC").Limit(limit).Find(&accounts)
	respondJSON(w, http.StatusOK, accounts)
}

// util function to retrieve models or 404
func getAccountOr404(db *gorm.DB, accountid string, w http.ResponseWriter, r *http.Request) *model.Account {
	account := model.Account{}
	if err := db.First(&account, model.Account{Aid: accountid}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &account
}
