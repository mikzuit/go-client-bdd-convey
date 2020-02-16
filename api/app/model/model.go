package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Account struct{

}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Account{})
	return db
}
