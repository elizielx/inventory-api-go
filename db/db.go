package db

import (
	"fmt"
	"github.com/elizielx/inventory-api-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDatabase(configuration config.Configuration) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", configuration.DatabaseHost, configuration.DatabaseUser, configuration.DatabasePassword, configuration.DatabaseName, configuration.DatabasePort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	permissionLevel := db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permission_level') THEN CREATE TYPE permission_level AS ENUM ('USER', 'ADMINISTRATOR', 'SUPERADMINISTRATOR'); END IF; END $$;")
	if permissionLevel.Error != nil {
		log.Fatal(permissionLevel.Error)
	}

	log.Println("Database successfully connected")
}

func GetDatabase() *gorm.DB {
	return db
}
