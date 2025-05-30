package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/m4d4r4666/fiber-crud/db"
	"github.com/m4d4r4666/fiber-crud/routes"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("error connecting DB: %v", err)
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(limiter.New())

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
