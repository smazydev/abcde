package utils

import (
	"log"

	"github.com/smazydev/abcde/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "postgres://postgres:C!tynet23@localhost:5432/erp"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Business{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate the database schema:", err)
	}
}
