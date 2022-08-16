package routes

import (
	"github.com/Guleri24/go-userapi-rest/controllers"
	"github.com/julienschmidt/httprouter"
)

func Routes(router *httprouter.Router) {
	router.GET("/", controllers.Welcome)
	router.GET("/api/v1/user/:id", controllers.GetUser)
	router.GET("/api/v1/user", controllers.GetUsers)
	router.POST("/api/v1/user", controllers.CreateUser)
	router.DELETE("/api/v1/user/:id", controllers.DeleteUser)
	router.PATCH("/api/v1/user", controllers.UpdateUser)

}
