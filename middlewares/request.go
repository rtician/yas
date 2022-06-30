package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitMiddlewares(app *fiber.App) {
	app.Use(recover.New())
	app.Use(JSONRequestType)
}

func JSONRequestType(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	return ctx.Next()
}
