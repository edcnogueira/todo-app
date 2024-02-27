package handlers

import "github.com/gofiber/fiber/v3"

var (
	fromProtected bool = false
)

func Setup(app *fiber.App) {
	app.Get("/", HandleViewHome)
	app.Get("/login", HandleViewLogin)
}
