package main

import (
	"github.com/elizielx/inventory-api-go/config"
	"github.com/elizielx/inventory-api-go/controllers"
	"github.com/elizielx/inventory-api-go/db"
	"github.com/elizielx/inventory-api-go/routes"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	server *gin.Engine

	UserController      *controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	db.InitDatabase(configuration)

	UserController = controllers.NewUserController(db.GetDatabase())
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	router := server.Group("/api")
	UserRouteController.UserRoute(router)

	log.Fatal(server.Run(":" + configuration.ServerPort))
}
