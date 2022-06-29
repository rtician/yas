package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"yas/cfg"
)

func Serve() error {
	app := fiber.New()

	log.Info().Msgf("starting listening on port %s", cfg.HttpServerPort())
	app.Listen(cfg.HttpServerPort())

	return nil
}
