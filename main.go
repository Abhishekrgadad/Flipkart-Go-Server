package main

import (
	"fmt"
	"log"
	"server/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.Load()

	db, err := config.ConnectDB(cfg.MONGO_URI)
	if err != nil {
		log.Fatal("failed to connect to DB")
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server is running successfully")
	})

	fmt.Println(cfg.PORT)
	fmt.Print(db.Name())
	log.Fatal(app.Listen(":8000"))
}
