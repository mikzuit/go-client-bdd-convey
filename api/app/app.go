package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mikzuit/api/config"
	"github.com/mikzuit/api/app/model"
	"github.com/mikzuit/api/app/handler"
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

	// Set new Routes
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters(){
	a.Post("/v1/organisation/accounts", a.CreateAccount)
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// indicated handler for create account
func (a *App) CreateAccount(w http.ResponseWriter, r *http.Request) {
	handler.CreateAccount(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
