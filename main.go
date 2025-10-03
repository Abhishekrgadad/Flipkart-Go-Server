package main

import (
	"log"
	"server/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()
	db := config.connectDB(cfg.MongoURI)

	app := fiber.New()

	app.use(fiber.logger())

	users.RegisterRoutes(app, db, cfg)

	log.printf("listening to port %s", cfg.port)
	log.Fatal(app.Listen(cfg.port))
}
