package endpoint

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler returns a generic error as text.
func ErrorHandler(c *gin.Context) {
	c.String(http.StatusInternalServerError, errors.New("generic error").Error())
}
