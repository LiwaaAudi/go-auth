package main

import (
	"github.com/LiwaaAudi/go-auth/repository"
	"github.com/LiwaaAudi/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository.Connect()

	app := fiber.New()
	
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	app.Listen(":8000")
}
