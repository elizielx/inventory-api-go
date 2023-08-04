package main

import (
	"github.com/elizielx/inventory-api-go/config"
	"github.com/elizielx/inventory-api-go/database"
	"github.com/elizielx/inventory-api-go/models"
	"log"
)

func init() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	database.Connect(configuration)
}

func main() {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	log.Println("Migration completed")
}
