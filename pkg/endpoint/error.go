package endpoint

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler returns a generic error as text.
func ErrorHandler(c *fiber.Ctx) error {
	return errors.New("generic error")
}
