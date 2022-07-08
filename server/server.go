package server

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"yas/cfg"
	"yas/database"
	"yas/handlers"
	"yas/middlewares"
	"yas/repositories"
	"yas/routes"
	"yas/services"
)

func Serve() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	mongoClient, err := database.GetMongoClient(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating mongodb connection")
		return err
	}

	database := mongoClient.Database(cfg.DbName())
	companyRepository := repositories.NewCompanyRepository(database)
	userRepository := repositories.NewUserRepository(database)

	companyService := services.NewCompanyService(companyRepository)
	userService := services.NewUserService(userRepository)

	apiHandlers := handlers.NewAPIHandlers(companyService)
	loginHandlers := handlers.NewLoginHandlers(userService)

	app := fiber.New()
	middlewares.InitMiddlewares(app)
	routes.LoginRoutes(app, loginHandlers)
	routes.APIRoutes(app, apiHandlers)

	log.Info().Msgf("starting listening on port %s", cfg.HttpServerPort())
	app.Listen(cfg.HttpServerPort())

	return nil
}
