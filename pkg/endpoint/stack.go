package endpoint

import "github.com/gofiber/fiber/v2"

// StackHandler returns the application stack as JSON.
func StackHandler(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}
