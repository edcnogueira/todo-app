package handlers

import (
	"github.com/a-h/templ"
	"github.com/edcnogueira/todo-app/views"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func HandleViewHome(ctx fiber.Ctx) error {
	hindex := views.HomeIndex(fromProtected)
	home := views.Home("", fromProtected, hindex)

	handler := adaptor.HTTPHandler(templ.Handler(home))

	return handler(ctx)
}
