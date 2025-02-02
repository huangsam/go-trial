package endpoint

import "github.com/gofiber/fiber/v2"

// HelloHandler returns a simple greeting as text.
func HelloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}
