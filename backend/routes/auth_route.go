package routes

import (
	"ecommerce-project/daos"
	"ecommerce-project/handlers"
	"ecommerce-project/services"
	"ecommerce-project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router) {
	authDAO := daos.NewUserDAO()
	authService := services.NewAuthService(authDAO)

	validate := validator.New()
	authHandler := handlers.NewAuthHandler(authService, &utils.Validator{Validator: validate})

	authGroup := app.Group("/auth")
	authGroup.Post("/", authHandler.Login)
}