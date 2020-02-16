package handler

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
	"bytes"
	"fmt"
	"github.com/mikzuit/api/app"
	"github.com/mikzuit/api/config"
	"github.com/mikzuit/api/app/model"
)

func TestCreateAccountEdward(t *testing.T) {
	var jsonString = []byte(`{
		"data": { 
			"type":"accounts",
			"id":"nd27e265-9605-4b4b-a0e5-3003ea9cc4di",
			"organization_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"version":0, 
			"attributes": {
				"country":"GB",
				"base_currency": "GBP",
				"account_number": "41426819",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22",
				"iban": "GB11NWBK40030041426819",
				"title": "Ms",
				"first_name": "Edward",
				"bank_account_name": "Edward Wilde",
				"alternative_bank_account_names": ["Edward Wilde"],
				"account_classification": "Personal",
				"joint_account": false,
				"account_matching_opt_out": false,
				"secondary_identification": "A1B2C3D4"
			}
		}
	}`)

	Convey("Given a POST request to /v1/organisation/accounts", t, func() {
		req := httptest.NewRequest("POST", "/v1/organisation/accounts", bytes.NewBuffer(jsonString))
		res := httptest.NewRecorder()
		Convey("When the request is handled by router", func() {
			config := config.GetConfig()
			app := &app.App{}
			app.Initialize(config)
			app.Router.ServeHTTP(res,req)
			Convey("then the response should be 201", func() {
				So(res.Code, ShouldEqual, 201)		
			})
		})
	})
}

func TestCreateAccountAlex(t *testing.T) {
	var jsonString = []byte(`{
		"data": { 
			"type":"accounts",
			"id":"vd45e265-1605-4b4b-a0e5-3003ea9cc6dc",
			"organization_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"version":0, 
			"attributes": {
				"country":"GB",
				"base_currency": "GBP",
				"account_number": "51426813",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22",
				"iban": "GB11NWBK40030041412345",
				"title": "Ms",
				"first_name": "Alex",
				"bank_account_name": "Alex Forsberg",
				"alternative_bank_account_names": ["Alex Forsberg"],
				"account_classification": "Personal",
				"joint_account": false,
				"account_matching_opt_out": false,
				"secondary_identification": "A1B2C3D4"
			}
		}
	}`)

	Convey("Given a POST request to /v1/organisation/accounts", t, func() {
		req := httptest.NewRequest("POST", "/v1/organisation/accounts", bytes.NewBuffer(jsonString))
		res := httptest.NewRecorder()
		Convey("When the request is handled by router", func() {
			config := config.GetConfig()
			app := &app.App{}
			app.Initialize(config)
			app.Router.ServeHTTP(res,req)
			Convey("then the response should be 201", func() {
				account := model.Account{}
				So(res.Code, ShouldEqual, 201)
				fmt.Printf("this models %v", account)
			})
		})
	})
}

func TestCreateAccountJames(t *testing.T) {
	var jsonString = []byte(`{
		"data": { 
			"type":"accounts",
			"id":"td27e265-9605-4b4b-a0e5-4004ea9cc4dm",
			"organization_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"version":0, 
			"attributes": {
				"country":"GB",
				"base_currency": "GBP",
				"account_number": "41426819",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22",
				"iban": "GB11NWBK40030041426819",
				"title": "Ms",
				"first_name": "James",
				"bank_account_name": "James Smith",
				"alternative_bank_account_names": ["James Smith"],
				"account_classification": "Personal",
				"joint_account": false,
				"account_matching_opt_out": false,
				"secondary_identification": "A1B2C3D4"
			}
		}
	}`)

	Convey("Given a POST request to /v1/organisation/accounts", t, func() {
		req := httptest.NewRequest("POST", "/v1/organisation/accounts", bytes.NewBuffer(jsonString))
		res := httptest.NewRecorder()
		Convey("When the request is handled by router", func() {
			config := config.GetConfig()
			app := &app.App{}
			app.Initialize(config)
			app.Router.ServeHTTP(res,req)
			Convey("then the response should be 201", func() {
				So(res.Code, ShouldEqual, 201)		
			})
		})
	})
}

func TestCreateAccountKevin(t *testing.T) {
	var jsonString = []byte(`{
		"data": { 
			"type":"accounts",
			"id":"be27e267-9608-4b4g-h0e5-4003ea9cc5fd",
			"organization_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"version":0, 
			"attributes": {
				"country":"GB",
				"base_currency": "GBP",
				"account_number": "51426810",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22",
				"iban": "GB11NWBK30030041434567",
				"title": "Ms",
				"first_name": "Kevin",
				"bank_account_name": "William Holditch",
				"alternative_bank_account_names": ["Kevin Holditch"],
				"account_classification": "Personal",
				"joint_account": false,
				"account_matching_opt_out": false,
				"secondary_identification": "A2B3C3D3"
			}
		}
	}`)

	Convey("Given a POST request to /v1/organisation/accounts", t, func() {
		req := httptest.NewRequest("POST", "/v1/organisation/accounts", bytes.NewBuffer(jsonString))
		res := httptest.NewRecorder()
		Convey("When the request is handled by router", func() {
			config := config.GetConfig()
			app := &app.App{}
			app.Initialize(config)
			app.Router.ServeHTTP(res,req)
			Convey("then the response should be 201", func() {
				So(res.Code, ShouldEqual, 201)		
			})
		})
	})
}

func TestPagination(t *testing.T){
	Convey("Given a GET request to /v1/organisation/accounts/2/2", t, func() {
		req := httptest.NewRequest("GET", "/v1/organisation/accounts/2/2", nil)
		res := httptest.NewRecorder()
		Convey("When the request is handler by router", func() {
			config := config.GetConfig()
			app := &app.App{}
			app.Initialize(config)
			app.Router.ServeHTTP(res,req)
			Convey("Then the response should be 200", func() {
				So(res.Code, ShouldEqual, 200)
			})
		})
	})
}

func TestFetchAccount(t *testing.T){
	Convey("Given a GET request to /v1/organisation/accounts/vd45e265-1605-4b4b-a0e5-3003ea9cc6dc", t, func() {
		req := httptest.NewRequest("GET", "/v1/organisation/accounts/vd45e265-1605-4b4b-a0e5-3003ea9cc6dc", nil)
		res := httptest.NewRecorder()
		Convey("When the request is handler by router", func() {
			config := config.GetConfig()
			app := &app.App{}
			app.Initialize(config)
			app.Router.ServeHTTP(res,req)
			Convey("Then the response should be 200", func() {
				So(res.Code, ShouldEqual, 200)
			})
		})
	})
}

func TestDeleteAccount(t *testing.T) {
	Convey("Given a DELETE request to /v1/organisation/accounts/td27e265-9605-4b4b-a0e5-4004ea9cc4dm?version=0", t, func() {
		req := httptest.NewRequest("DELETE", "/v1/organisation/accounts/td27e265-9605-4b4b-a0e5-4004ea9cc4dm?version=0", nil)
		res := httptest.NewRecorder()
		Convey("When the request is handler by router", func() {
			config := config.GetConfig()
			app := &app.App{}
			app.Initialize(config)
			app.Router.ServeHTTP(res,req)
			Convey("Then the response should be 204", func() {
				So(res.Code, ShouldEqual, 204)
			})
		})
	})
}
