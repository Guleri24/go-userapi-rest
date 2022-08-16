package main

import (
	"net/http"

	"github.com/Guleri24/go-userapi-rest/database"
	"github.com/Guleri24/go-userapi-rest/routes"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	database.MongodbAtlasSetup()
	routes.Routes(router)
	http.ListenAndServe(":6969", router)
}
