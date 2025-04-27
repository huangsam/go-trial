package endpoint_test

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/huangsam/go-trial/internal/model"
	"github.com/huangsam/go-trial/internal/util"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	router := echo.New()

	testCases := []struct {
		name             string
		path             string
		query            string
		handler          echo.HandlerFunc
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
			expectedContains: []string{"generic", "error"},
		},
		{
			name:             "RectangleWithQuery",
			path:             "/rectangle-size",
			query:            "?width=3.14&height=3.14",
			handler:          endpoint.RectangleSizeHandler,
			expectedStatus:   http.StatusOK,
			expectedContains: []string{"width", "height"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router.GET(tc.path, tc.handler)
			req := httptest.NewRequest(http.MethodGet, tc.path+tc.query, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			resp := w.Result()
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

func TestBasicAuthHandler(t *testing.T) {
	router := echo.New()
	acc := model.UserAccount{Username: "foo", Password: "bar"}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(acc.Username + ":" + acc.Password))

	router.GET("/secret", endpoint.HelloHandler, util.SetupBasicAuth(acc))

	t.Run("NoAuthHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/secret", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, string(body), "Body should not be empty")
		assert.Contains(t, string(body), "Unauthorized")
	})

	t.Run("ValidAuthHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/secret", nil)
		req.Header.Set("Authorization", "Basic "+basicAuth)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, string(body), "Body should not be empty")
		assert.Contains(t, string(body), "Hello")
	})
}
