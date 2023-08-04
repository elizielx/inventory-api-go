package main

import (
	"github.com/elizielx/inventory-api-go/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine
)

func main() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	server = gin.Default()
	router := server.Group("/api")

	router.GET("/health", func(ctx *gin.Context) {
		message := "OK"
		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	log.Fatal(server.Run(":" + configuration.ServerPort))
}
