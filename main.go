package main

import (
	"github.com/Guleri24/go-userapi-rest/configs"
	_ "github.com/Guleri24/go-userapi-rest/docs"
	"github.com/Guleri24/go-userapi-rest/routes"
	"github.com/gofiber/fiber/v2"
)

// @title User Information RESTful API on Go: Fiber, MongoDB Atlas
// @version 1.0
// @description REST API for a User Information application in which we create new users, view them, update & delete their information.

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host go-userapi-rest.herokuapp.com
// @BasePath /api/v1

func main() {
	// Define a new Fiber
	app := fiber.New()

	// Database.
	configs.ConnectDB() // Register a connection to Database using .Env Configs

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.UserRoute(app)     // Register routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	app.Listen(":" + configs.Port())
}
