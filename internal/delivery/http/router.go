package http

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

}
