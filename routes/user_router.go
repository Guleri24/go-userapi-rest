package routes

import (
	"github.com/Guleri24/go-userapi-rest/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// Create routes group.
	route := app.Group("/api/v1")

	// Routes for Post method:
	route.Post("/user", controllers.CreateUser) // create a new user

	// Routes for Get method:
	route.Get("/user/:id", controllers.GetAUser) // get one user by ID
	route.Get("/users", controllers.GetAllUsers) // get all the users

	// Routes for Patch method:
	route.Patch("/user/:id", controllers.EditAUser) // update one user by ID

	//Routes for Delete method:
	route.Delete("/user/:id", controllers.DeleteAUser) // delete one user by ID
}
