package advislistdata

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
	app.DB().RegisterTable("main", AdvisListData{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&AdvisListData{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"advis_list_data.detail",
		"advis_list_data.list",
		"advis_list_data.create",
		"advis_list_data.edit",
		"advis_list_data.delete",
	}))
	app.Server().AddRoute("/advis_list_data", "POST", REST().Create, nil)
	app.Server().AddRoute("/advis_list_data", "GET", REST().Get, nil)
	app.Server().AddRoute("/advis_list_data/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/advis_list_data/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/advis_list_data/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/advis_list_data/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestAdvisListDataID returns an available AdvisListData ID.
func getTestAdvisListDataID() string {
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
		description:  "Get empty list of AdvisListData",
		method:       "GET",
		path:         "/advis_list_data",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create AdvisListData with minimum payload",
		method:       "POST",
		path:         "/advis_list_data",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get AdvisListData by ID",
		method:       "GET",
		path:         "/advis_list_data/" + getTestAdvisListDataID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update AdvisListData by ID",
		method:       "PUT",
		path:         "/advis_list_data/" + getTestAdvisListDataID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update AdvisListData by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update AdvisListData by ID",
		method:       "PATCH",
		path:         "/advis_list_data/" + getTestAdvisListDataID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update AdvisListData by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete AdvisListData by ID",
		method:       "DELETE",
		path:         "/advis_list_data/" + getTestAdvisListDataID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete AdvisListData by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestAdvisListDataREST tests the REST API of AdvisListData data with specified scenario.
func TestAdvisListDataREST(t *testing.T) {
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

// BenchmarkAdvisListDataREST tests the REST API of AdvisListData data with specified scenario.
func BenchmarkAdvisListDataREST(b *testing.B) {
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
