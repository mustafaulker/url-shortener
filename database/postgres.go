package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"url-shortener/internal/models"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.ShortURL{}); err != nil {
		return nil, err
	}
	return db, nil
}
