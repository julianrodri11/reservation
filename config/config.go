package config

import (
	"fmt"
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

	// Llamar a la función DropTables para eliminar las tablas
	/* if err := DropTables(db); err != nil {
		log.Fatal(err)
	} */

	// Migrar las entidades nuevamente
	err = db.AutoMigrate(&entity.Users{}, &entity.Reservations{})
	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	return db
}

func DropTables(db *gorm.DB) error {
	// Eliminar las tablas Users y Reservations
	err := db.Migrator().DropTable(&entity.Users{}, &entity.Reservations{})
	if err != nil {
		return fmt.Errorf("Failed to drop tables: %w", err)
	}
	return nil
}
