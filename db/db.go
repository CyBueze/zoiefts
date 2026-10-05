package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"zoie/models"
)

func Connect(databaseURL string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := conn.AutoMigrate(&models.Donation{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return conn
}