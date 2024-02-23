package main

import (
	"log"

	"github.com/edcnogueira/todo-app/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {

	app := fiber.New()
	app.Static("/", "./assets")

	app.Use(logger.New())

	handlers.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
