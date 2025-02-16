package util

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupBasicAuth sets up basic authentication middleware for an Echo web server.
//
// It uses the provided username and password to authenticate requests.
// Returns an Echo middleware function that checks the provided credentials against
// predefined admin credentials (AdminUser and AdminPass).
// If the credentials are valid, the request is allowed to proceed; otherwise, an error is returned.
func SetupBasicAuth(username, password string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		if u == username && p == password {
			return true, nil
		}
		return false, errors.New("invalid user credentials")
	})
}
