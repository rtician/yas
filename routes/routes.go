package routes

import (
	"yas/handlers"

	"github.com/gofiber/fiber/v2"
)

func APIRoutes(app *fiber.App, h handlers.APIHandler) {
	route := app.Group("/api")
	route.Post("/companies", h.CreateCompanyHandler)
}
