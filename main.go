package main

import (
	"github.com/LiwaaAudi/go-auth/repository"
	"github.com/LiwaaAudi/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	repository.Connect()

	app := fiber.New()

	routes.Setup(app)
	app.Listen(":8000")
}
