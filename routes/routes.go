package routes

import (
	"yas/handlers"
	"yas/middlewares"

	"github.com/gofiber/fiber/v2"
)

func APIRoutes(app *fiber.App, h handlers.APIHandler) {
	route := app.Group("/api")
	route.Post("/companies", middlewares.JWTAuth, h.CreateCompanyHandler)
}

func LoginRoutes(app *fiber.App, h handlers.LoginHandler) {
	app.Post("/register", h.RegisterHandler)
	app.Post("/login", h.LoginHandler)
}
