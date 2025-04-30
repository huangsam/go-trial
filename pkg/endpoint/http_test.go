package endpoint_test

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/huangsam/go-trial/internal/model"
	"github.com/huangsam/go-trial/internal/util"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	r := chi.NewRouter()

	testCases := []struct {
		name             string
		path             string
		query            string
		handler          http.HandlerFunc
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
			expectedContains: []string{"Generic", "error"},
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
			r.Get(tc.path, tc.handler)
			req := httptest.NewRequest(http.MethodGet, tc.path+tc.query, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			resp := w.Result()
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.NotEmpty(t, string(body), "Body should not be empty")
			responseBody := string(body)
			for _, substring := range tc.expectedContains {
				assert.Contains(t, responseBody, substring, "Body should contain %s", substring)
			}
		})
	}
}

func TestBasicAuthHandler(t *testing.T) {
	r := chi.NewRouter()
	acc := model.UserAccount{Username: "foo", Password: "bar"}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(acc.Username + ":" + acc.Password))

	r.With(util.BasicAuth(acc)).Get("/secret", endpoint.HelloHandler)

	t.Run("NoAuthHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/secret", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.NotEmpty(t, string(body), "Body should not be empty")
		assert.Contains(t, string(body), "Unauthorized")
	})

	t.Run("ValidAuthHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/secret", nil)
		req.Header.Set("Authorization", "Basic "+basicAuth)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.NotEmpty(t, string(body), "Body should not be empty")
		assert.Contains(t, string(body), "Hello")
	})
}
