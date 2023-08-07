package main

import (
	"github.com/elizielx/inventory-api-go/config"
	"github.com/elizielx/inventory-api-go/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine
)

func init() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	db.InitDatabase(configuration)
	server = gin.Default()
}

func main() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	router := server.Group("/api")
	router.GET("/health", func(ctx *gin.Context) {
		message := "OK"
		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	log.Fatal(server.Run(":" + configuration.ServerPort))
}
