package handler

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/mikzuit/api/app"
	"github.com/mikzuit/api/config"
)

func TestCreateAccount(t *testing.T) {
	var jsonString = []byte(`{
		"data": { 
			"type":"accounts",
			"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
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
				"first_name": "Samantha",
				"bank_account_name": "Samantha Holder",
				"alternative_bank_account_names": ["Sam Holder"],
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
