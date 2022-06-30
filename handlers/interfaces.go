package handlers

import (
	"yas/services"

	"github.com/gofiber/fiber/v2"
)

type APIHandler interface {
	CreateCompanyHandler(ctx *fiber.Ctx) error
}

type apiHandler struct {
	companyService services.CompanyService
}

func NewAPIHandlers(companyService services.CompanyService) APIHandler {
	return &apiHandler{
		companyService: companyService,
	}
}
