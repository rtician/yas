package handlers

import (
	"yas/services"
	"yas/types"
	"yas/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type loginHandler struct {
	userService services.UserService
}

func NewLoginHandlers(s services.UserService) LoginHandler {
	return &loginHandler{
		userService: s,
	}
}

func (h *loginHandler) RegisterHandler(ctx *fiber.Ctx) error {
	u := &types.UserRegistration{}
	if err := ctx.BodyParser(u); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	errors := utils.ValidateStruct(u)
	if errors != nil {
		return ctx.JSON(errors)
	}

	createdUser, err := h.userService.Register(ctx.Context(), u)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdUser)
}

func (h *loginHandler) LoginHandler(ctx *fiber.Ctx) error {
	login := &types.UserLogin{}
	if err := ctx.BodyParser(login); err != nil {
		log.Error().Err(err)
		return fiber.NewError(fiber.StatusBadGateway)
	}

	errors := utils.ValidateStruct(login)
	if errors != nil {
		return ctx.JSON(errors)
	}

	token, err := h.userService.Login(ctx.Context(), login)
	if err != nil {
		log.Error().Err(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
