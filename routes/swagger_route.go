package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//SwaggerRoute func for descibe group of API Docs routes.
func SwaggerRoute(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
