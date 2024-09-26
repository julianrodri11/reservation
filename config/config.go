package config

import (
	"log"
	"reservation-system/models/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=reservas port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrar las entidades
	db.AutoMigrate(&entity.Users{}, &entity.Reservations{})
	return db
}
