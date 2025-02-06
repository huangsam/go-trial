package util

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	AdminUser = "admin"
	AdminPass = "admin"
)

// SetupBasicAuth sets up basic authentication middleware for an Echo web server.
//
// It uses the provided username and password to authenticate requests.
// Returns an Echo middleware function that checks the provided credentials against
// predefined admin credentials (AdminUser and AdminPass).
// If the credentials are valid, the request is allowed to proceed; otherwise, an error is returned.
func SetupBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == AdminUser && password == AdminPass {
			return true, nil
		}
		return false, errors.New("invalid user credentials")
	})
}
