package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mikzuit/api/config"
	"github.com/mikzuit/api/app/model"
)

// App Router and DB instances

type App struct {
	Router *mux.Router
	DB *gorm.DB
}

// App initializator with predefined config
func(a *App) Initialize(config *config.Config) {

	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.DB.Host,
	5432,
	config.DB.Username,
	config.DB.Password,
	config.DB.Name)	

	// connection to database
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	// database population models
	a.DB = model.DBMigrate(db)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
