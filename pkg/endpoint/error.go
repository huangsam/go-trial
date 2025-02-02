package endpoint

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler returns a generic error.
func ErrorHandler(c *fiber.Ctx) error {
	return errors.New("What is going on with the world?")
}
