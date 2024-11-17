package config

import (
	"fmt"
	"log"
	"os"
	"reservation-system/models/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	// Obtener variables de entorno
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Llamar a la función DropTables para eliminar las tablas
	/* if err := DropTables(db); err != nil {
		log.Fatal(err)
	} */

	// Migrar las entidades nuevamente
	err = db.AutoMigrate(&entity.Company{}, &entity.Users{}, &entity.Reservations{})
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
