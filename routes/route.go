package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jxtsly111/blog/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
