package routes

import (
	"ecommerce-project/daos"
	"ecommerce-project/handlers"
	"ecommerce-project/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	userDAO := daos.NewUserDAO()
	userService := services.NewUserService(userDAO)
	userHandler := handlers.NewUserHandler(userService)

	userGroup := app.Group("/users")
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Get("/", userHandler.GetUser)
}