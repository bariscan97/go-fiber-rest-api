package main

import (
	"go-pgx/interval/middleware"
	"go-pgx/interval/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	app.Use(logger.New())
	app.Use(middleware.Auth)
	routes.Routers(app)
	log.Fatal(app.Listen(os.Getenv("PORT")))

}
