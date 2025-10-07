package tests

import (
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"github.com/pretorian41/goaggregate/server"
	"github.com/stretchr/testify/assert"
)

var app = server.New()

func TestRoutes(t *testing.T) {
	tests := []struct {
		testName     string
		route        string
		method       string
		body         io.Reader
		expectedCode int
		expectedBody string
	}{
		{
			testName:     "an invalid path",
			route:        "/api/agg/invalid-path",
			method:       "GET",
			body:         nil,
			expectedCode: 404,
			expectedBody: "{\"error\":\"Cannot GET /api/agg/invalid-path\"}",
		},
		{
			testName:     "another invalid path",
			route:        "/api/agg/another-invalid-path",
			method:       "POST",
			body:         nil,
			expectedCode: 404,
			expectedBody: "{\"error\":\"Cannot POST /api/agg/another-invalid-path\"}",
		},
		{
			testName:     "a valid path",
			route:        "/api/agg/fetch/test",
			method:       "GET",
			body:         nil,
			expectedBody: "{\"avatar_url\":\"https://i.pravatar.cc/300\",\"email\":\"test@test.com\",\"id\":\"test\",\"name\":\"John Foo\"}",
			expectedCode: 200,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			test.body,
		)

		res, err := app.Test(req, -1)
		assert.Nilf(t, err, test.testName)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.testName)

		if test.expectedBody != "" {
			body, err := ioutil.ReadAll(res.Body)
			assert.Nilf(t, err, test.testName)
			assert.Equalf(t, test.expectedBody, string(body), test.testName)
		}

	}
}
