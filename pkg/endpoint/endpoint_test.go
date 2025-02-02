package endpoint_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/stretchr/testify/assert"
)

func TestHandlerStatus(t *testing.T) {
	app := fiber.New()

	testCases := []struct {
		name             string
		path             string
		query            string
		handler          fiber.Handler
		expectedStatus   int
		expectedContains []string
	}{
		{
			name:             "Hello",
			path:             "/",
			handler:          endpoint.HelloHandler,
			expectedStatus:   http.StatusOK,
			expectedContains: []string{"Hello"},
		},
		{
			name:             "Error",
			path:             "/error",
			handler:          endpoint.ErrorHandler,
			expectedStatus:   http.StatusInternalServerError,
			expectedContains: []string{"What", "the world?"},
		},
		{
			name:             "Rectangle default",
			path:             "/rectangle-size",
			handler:          endpoint.RectangleSizeHandler,
			expectedStatus:   http.StatusOK,
			expectedContains: []string{"width", "height"},
		},
		{
			name:             "Rectangle with query",
			path:             "/rectangle-size",
			query:            "?width=3.14&height=3.14",
			handler:          endpoint.RectangleSizeHandler,
			expectedStatus:   http.StatusOK,
			expectedContains: []string{"width", "height"},
		},
		{
			name:             "Stack",
			path:             "/stack",
			handler:          endpoint.StackHandler,
			expectedStatus:   http.StatusOK,
			expectedContains: []string{"GET", "HEAD"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app.Get(tc.path, tc.handler)
			req := httptest.NewRequest(http.MethodGet, tc.path+tc.query, nil)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, string(body), "Body should not be empty")
			responseBody := string(body)
			for _, substring := range tc.expectedContains {
				assert.True(t, strings.Contains(responseBody, substring), "Body should contain %s", substring)
			}
		})
	}
}
