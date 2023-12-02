package advisanalytic

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
	app.DB().RegisterTable("main", AdvisAnalytic{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&AdvisAnalytic{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"advis_analytics.detail",
		"advis_analytics.list",
		"advis_analytics.create",
		"advis_analytics.edit",
		"advis_analytics.delete",
	}))
	app.Server().AddRoute("/advis_analytics", "POST", REST().Create, nil)
	app.Server().AddRoute("/advis_analytics", "GET", REST().Get, nil)
	app.Server().AddRoute("/advis_analytics/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/advis_analytics/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/advis_analytics/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/advis_analytics/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestAdvisAnalyticID returns an available AdvisAnalytic ID.
func getTestAdvisAnalyticID() string {
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
		description:  "Get empty list of AdvisAnalytic",
		method:       "GET",
		path:         "/advis_analytics",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create AdvisAnalytic with minimum payload",
		method:       "POST",
		path:         "/advis_analytics",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get AdvisAnalytic by ID",
		method:       "GET",
		path:         "/advis_analytics/" + getTestAdvisAnalyticID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update AdvisAnalytic by ID",
		method:       "PUT",
		path:         "/advis_analytics/" + getTestAdvisAnalyticID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update AdvisAnalytic by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update AdvisAnalytic by ID",
		method:       "PATCH",
		path:         "/advis_analytics/" + getTestAdvisAnalyticID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update AdvisAnalytic by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete AdvisAnalytic by ID",
		method:       "DELETE",
		path:         "/advis_analytics/" + getTestAdvisAnalyticID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete AdvisAnalytic by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestAdvisAnalyticREST tests the REST API of AdvisAnalytic data with specified scenario.
func TestAdvisAnalyticREST(t *testing.T) {
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

// BenchmarkAdvisAnalyticREST tests the REST API of AdvisAnalytic data with specified scenario.
func BenchmarkAdvisAnalyticREST(b *testing.B) {
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
