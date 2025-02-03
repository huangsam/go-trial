package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler returns a simple greeting as text.
func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}
