package routes

import (
	"Github.com/Guleri24/userapi/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	route := app.Group("/api/v1")
	route.Post("/user", controllers.CreateUser)
	route.Get("/user/:userId", controllers.GetAUser)
}
