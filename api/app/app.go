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

	// Connection to database
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	// Database population models
	a.DB = model.DBMigrate(db)

	// Set new Routes
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters(){
	a.Get("/v1/organisation/accounts/{page_number}/{page_size}", a.GetPagination)
	a.Get("/v1/organisation/accounts/{accountid}", a.FetchAccount)
	//a.Get("/v1/organisation/accounts?pagenumber={page_number}&pagesize={page_size}", a.GetPagination) //&filter[{attribute}]={filter_value}
	a.Post("/v1/organisation/accounts", a.CreateAccount)
	a.Delete("/v1/organisation/accounts/{accountid}", a.DeleteAccount)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Indicated handler for create account
func (a *App) CreateAccount(w http.ResponseWriter, r *http.Request) {
	handler.CreateAccount(a.DB, w, r)
}

// Indicated handler for fetch account
func (a *App) FetchAccount(w http.ResponseWriter, r *http.Request) {
	handler.FetchAccount(a.DB, w, r)
}

// Indicated handler for delete account
func (a *App) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	handler.DeleteAccount(a.DB, w, r)
}

// Indicated handler for account with pagination
func (a *App) GetPagination(w http.ResponseWriter, r *http.Request) {
	handler.GetPagination(a.DB, w, r)
}

// Run app on router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
