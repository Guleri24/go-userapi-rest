package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(app *fiber.App) {
	// Register new special route.
	app.Use(
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON reponse.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": http.StatusNotFound,
				"error":  true,
				"msg":    "sorry, endpoint is not found",
			})
		},
	)
}
