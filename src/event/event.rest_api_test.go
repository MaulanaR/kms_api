package event

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
	app.DB().RegisterTable("main", Event{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Event{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"events.detail",
		"events.list",
		"events.create",
		"events.edit",
		"events.delete",
	}))
	app.Server().AddRoute("/events", "POST", REST().Create, nil)
	app.Server().AddRoute("/events", "GET", REST().Get, nil)
	app.Server().AddRoute("/events/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/events/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/events/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/events/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestEventID returns an available Event ID.
func getTestEventID() string {
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
		description:  "Get empty list of Event",
		method:       "GET",
		path:         "/events",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create Event with minimum payload",
		method:       "POST",
		path:         "/events",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get Event by ID",
		method:       "GET",
		path:         "/events/" + getTestEventID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update Event by ID",
		method:       "PUT",
		path:         "/events/" + getTestEventID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update Event by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update Event by ID",
		method:       "PATCH",
		path:         "/events/" + getTestEventID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update Event by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete Event by ID",
		method:       "DELETE",
		path:         "/events/" + getTestEventID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete Event by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestEventREST tests the REST API of Event data with specified scenario.
func TestEventREST(t *testing.T) {
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

// BenchmarkEventREST tests the REST API of Event data with specified scenario.
func BenchmarkEventREST(b *testing.B) {
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
