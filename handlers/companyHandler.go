package handlers

import (
	"yas/database/models"
	"yas/utils"

	"github.com/gofiber/fiber/v2"
)

// CreateCompanyHandler func create a new company
// @Description Create a new company
// @Summary create new companies
// @Tags Company
// @Accepts json
// @Produces json
// @Success 201 {object} models.Company
// @Router /companies [post]
func (h *apiHandler) CreateCompanyHandler(ctx *fiber.Ctx) error {
	c := models.NewCompany()
	if err := ctx.BodyParser(c); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	errors := utils.ValidateStruct(c)
	if errors != nil {
		return ctx.JSON(errors)
	}

	createdCompany, err := h.companyService.CreateCompany(ctx.Context(), c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(createdCompany)
}
