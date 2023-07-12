package akademi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"

	"github.com/maulanar/kms/app"
)

// prepareTest prepares the test.
func prepareTest(tb testing.TB) {
	app.Test()
	tx := app.Test().Tx
	app.DB().RegisterTable("main", Akademi{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Akademi{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"akademi.detail",
		"akademi.list",
		"akademi.create",
		"akademi.edit",
		"akademi.delete",
	}))
	app.Server().AddRoute("/akademi", "POST", REST().Create, nil)
	app.Server().AddRoute("/akademi", "GET", REST().Get, nil)
	app.Server().AddRoute("/akademi/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/akademi/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/akademi/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/akademi/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestAkademiID returns an available Akademi ID.
func getTestAkademiID() string {
	return "todo"
}

// tests is test scenario.
var tests = []struct {
	description  string // description of the test case
	method       string // method to test
	path         string // route path to test
	token        string // token to test
	bodyRequest  string // body to test
	expectedCode int    // expected HTTP status code
	expectedBody string // expected body response
}{
	{
		description:  "Get empty list of Akademi",
		method:       "GET",
		path:         "/akademi",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create Akademi with minimum payload",
		method:       "POST",
		path:         "/akademi",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get Akademi by ID",
		method:       "GET",
		path:         "/akademi/" + getTestAkademiID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update Akademi by ID",
		method:       "PUT",
		path:         "/akademi/" + getTestAkademiID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update Akademi by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update Akademi by ID",
		method:       "PATCH",
		path:         "/akademi/" + getTestAkademiID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update Akademi by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete Akademi by ID",
		method:       "DELETE",
		path:         "/akademi/" + getTestAkademiID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete Akademi by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestAkademiREST tests the REST API of Akademi data with specified scenario.
func TestAkademiREST(t *testing.T) {
	prepareTest(t)

	// Iterate through test single test cases
	for _, test := range tests {

		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
		req.Header.Add("Authorization", "Bearer "+test.token)
		req.Header.Add("Content-Type", "application/json")

		// Perform the request plain with the app, the second argument is a request latency (set to -1 for no latency)
		res, err := app.Server().Test(req)

		// Verify if the status code is as expected
		utils.AssertEqual(t, nil, err, "app.Server().Test(req)")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)

		// Verify if the body response is as expected
		body, err := io.ReadAll(res.Body)
		utils.AssertEqual(t, nil, err, "io.ReadAll(res.Body)")
		app.Test().AssertMatchJSONElement(t, []byte(test.expectedBody), body, test.description)
		res.Body.Close()
	}
}

// BenchmarkAkademiREST tests the REST API of Akademi data with specified scenario.
func BenchmarkAkademiREST(b *testing.B) {
	b.ReportAllocs()
	prepareTest(b)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
			req.Header.Add("Authorization", "Bearer "+test.token)
			req.Header.Add("Content-Type", "application/json")
			app.Server().Test(req)
		}
	}
}
