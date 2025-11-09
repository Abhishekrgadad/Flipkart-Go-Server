package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() *config {

	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	mongouri := os.Getenv("MONGO_URI")
	if mongouri == "" {
		log.Fatal("mongourl is required")
	}

	jwtsecret := os.Getenv("JWT_SECRET")
	if jwtsecret == "" {
		log.Fatal("Jwt key required")
	}

	return &config{
		PORT:      port,
		MONGO_URI: mongouri,
		JWT_SECRET: jwtsecret,
	}
}
