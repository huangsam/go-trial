package endpoint

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorHandler returns a generic error as text.
func ErrorHandler(c echo.Context) error {
	return c.String(http.StatusInternalServerError, errors.New("generic error").Error())
}
