package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	env := os.Getenv("ENV")
	var err error

	// Cargar el archivo correspondiente según el entorno
	if env == "docker" {
		err = godotenv.Load(".env.docker")
	} else {
		err = godotenv.Load(".env.local")
	}

	if err != nil {
		log.Fatal("Error loading env file")
	}
}
