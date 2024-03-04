package notifikasi

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
	app.DB().RegisterTable("main", Notifikasi{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Notifikasi{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"notifikasi.detail",
		"notifikasi.list",
		"notifikasi.create",
		"notifikasi.edit",
		"notifikasi.delete",
	}))
	app.Server().AddRoute("/notifikasi", "POST", REST().Create, nil)
	app.Server().AddRoute("/notifikasi", "GET", REST().Get, nil)
	app.Server().AddRoute("/notifikasi/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/notifikasi/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/notifikasi/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/notifikasi/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestNotifikasiID returns an available Notifikasi ID.
func getTestNotifikasiID() string {
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
		description:  "Get empty list of Notifikasi",
		method:       "GET",
		path:         "/notifikasi",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create Notifikasi with minimum payload",
		method:       "POST",
		path:         "/notifikasi",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get Notifikasi by ID",
		method:       "GET",
		path:         "/notifikasi/" + getTestNotifikasiID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update Notifikasi by ID",
		method:       "PUT",
		path:         "/notifikasi/" + getTestNotifikasiID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update Notifikasi by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update Notifikasi by ID",
		method:       "PATCH",
		path:         "/notifikasi/" + getTestNotifikasiID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update Notifikasi by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete Notifikasi by ID",
		method:       "DELETE",
		path:         "/notifikasi/" + getTestNotifikasiID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete Notifikasi by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestNotifikasiREST tests the REST API of Notifikasi data with specified scenario.
func TestNotifikasiREST(t *testing.T) {
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

// BenchmarkNotifikasiREST tests the REST API of Notifikasi data with specified scenario.
func BenchmarkNotifikasiREST(b *testing.B) {
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