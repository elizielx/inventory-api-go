package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine
)

func main() {
	server = gin.Default()
	router := server.Group("/api")

	router.GET("/health", func(ctx *gin.Context) {
		message := "OK"
		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	log.Fatal(server.Run(":8080"))

}
