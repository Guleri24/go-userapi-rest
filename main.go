package main

import (
	"Github.com/Guleri24/userapi/configs"
	"Github.com/Guleri24/userapi/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}
