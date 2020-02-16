package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Record represents a record of data
type Payload struct {
	Data Account `json:"data"`
}

// Structs created with json-to-go website https://mholt.github.io/json-to-go/
type Attributes struct{
	Id string `gorm:"primary_key;unique_index"`
	Country string `gorm:"default:'GB'" json:"country"`
	BaseCurrency string `json:"base_currency"`
	AccountNumber string `json:"account_number"`
	BankId	string	`json:"bank_id"`
	BankIdCode string `json:"bank_id_code"`
	BIC string `json:"bic"`
	Iban string `json:"iban"`
	Title string `json:"title"`
	Firstname string `json:"first_name"`
	BankAccountName string `json:"bank_account_name"`
	AccountClassification string `json:"account_classification"`
	JointAccount bool `json:"joint_account"`
	AccountMatchingOptOut bool `json:"account_matching_opt_out"`
	SecondaryIdentification string `json:"secondary_identification"`
}

type Account struct{
	Type string `json:"type"`
	Aid string	`gorm:"primary_key" json:"id,omitempty"`
	OrganisationId string `json:"organisation_id"`
	Version int `json:"version"`
	Attributes Attributes `json:"attributes" gorm:"ForeignKey:Id"`
}


func DBMigrate(db *gorm.DB) *gorm.DB {	func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Account{})
	return db
}
