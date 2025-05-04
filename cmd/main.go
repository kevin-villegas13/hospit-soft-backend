package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"hospit-soft-backend/internal/delivery/http"
	"hospit-soft-backend/internal/infrastructure/database"
)

func main() {
	_ = godotenv.Load(".env")

	database.Connect()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	http.Router(app)

	log.Fatal(app.Listen(":3000"))
}
