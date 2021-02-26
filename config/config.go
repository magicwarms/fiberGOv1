package config

import (
	"github.com/gofiber/fiber/v2"
)

// NotFoundConfig is to handle route 404 not found exception
func NotFoundConfig(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}
