package main

import (
	"github.com/elizielx/inventory-api-go/config"
	"github.com/elizielx/inventory-api-go/db"
	"github.com/elizielx/inventory-api-go/models"
	"log"
)

func init() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	db.InitDatabase(configuration)
}

func main() {
	err := db.GetDatabase().AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	log.Println("Migration completed")
}
