package endpoint_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/huangsam/go-trial/lesson/endpoint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawHandlers(t *testing.T) {
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

func TestJsonHandlers(t *testing.T) {
	r := chi.NewRouter()

	testCases := []struct {
		name           string
		path           string
		query          string
		handler        http.HandlerFunc
		expectedStatus int
		expectedSize   string
	}{
		{
			name:           "RectangleWithQuery",
			path:           "/rectangle-size",
			query:          "?width=1&height=1",
			handler:        endpoint.RectangleSizeHandler,
			expectedStatus: http.StatusOK,
			expectedSize:   "Small",
		},
		{
			name:           "RectangleWithoutQuery",
			path:           "/rectangle-size",
			query:          "",
			handler:        endpoint.RectangleSizeHandler,
			expectedStatus: http.StatusOK,
			expectedSize:   "Small",
		},
		{
			name:           "CircleWithQuery",
			path:           "/circle-size",
			query:          "?radius=1",
			handler:        endpoint.CircleSizeHandler,
			expectedStatus: http.StatusOK,
			expectedSize:   "Small",
		},
		{
			name:           "CircleWithoutQuery",
			path:           "/circle-size",
			query:          "",
			handler:        endpoint.CircleSizeHandler,
			expectedStatus: http.StatusOK,
			expectedSize:   "Small",
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
			assert.NotEmpty(t, body, "Body should not be empty")

			var payload map[string]any
			require.NoError(t, json.Unmarshal(body, &payload), "Failed to unmarshal JSON response")
			assert.Greater(t, payload["area"], 0.0, "Area should be positive")
			assert.Greater(t, payload["perimeter"], 0.0, "Perimeter should be positive")
			assert.Equal(t, tc.expectedSize, payload["size"], "Size should be %s", tc.expectedSize)
		})
	}
}
