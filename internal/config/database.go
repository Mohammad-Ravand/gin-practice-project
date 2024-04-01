package config

import (
	models "example.com/gin-first-project/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Product{}, &models.User{})
	return db
}
