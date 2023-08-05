package database

import (
	"fmt"
	"github.com/elizielx/inventory-api-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect(config config.Configuration) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", config.DatabaseHost, config.DatabaseUser, config.DatabasePassword, config.DatabaseName, config.DatabasePort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	permissionLevel := DB.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permission_level') THEN CREATE TYPE permission_level AS ENUM ('USER', 'ADMINISTRATOR', 'SUPERADMINISTRATOR'); END IF; END $$;")
	if permissionLevel.Error != nil {
		log.Fatal(permissionLevel.Error)
	}

	log.Println("Database connected")
}
