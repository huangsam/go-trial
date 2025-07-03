package lesson_test

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/huangsam/go-trial/internal/lesson"
	"github.com/huangsam/go-trial/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicAuthMiddleware(t *testing.T) {
	r := chi.NewRouter()
	validAccount := model.UserAccount{Username: "testuser", Password: "testpass"}
	invalidAccount := model.UserAccount{Username: "invaliduser", Password: "invalidpass"}
	basicAuthHeader := base64.StdEncoding.EncodeToString([]byte(validAccount.Username + ":" + validAccount.Password))
	invalidAuthHeader := base64.StdEncoding.EncodeToString([]byte(invalidAccount.Username + ":" + invalidAccount.Password))

	r.With(lesson.BasicAuth(validAccount)).Get("/protected", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Protected content"))
	})

	t.Run("NoAuthHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.NotEmpty(t, string(body), "Body should not be empty")
		assert.Contains(t, string(body), "Unauthorized")
	})

	t.Run("InvalidAuthHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Basic "+invalidAuthHeader)
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
		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Basic "+basicAuthHeader)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.NotEmpty(t, string(body), "Body should not be empty")
		assert.Contains(t, string(body), "Protected content")
	})
}
