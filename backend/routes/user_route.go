package routes

import (
	"ecommerce-project/daos"
	"ecommerce-project/handlers"
	"ecommerce-project/middlewares"
	"ecommerce-project/services"
	"ecommerce-project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	userDAO := daos.NewUserDAO()
	userService := services.NewUserService(userDAO)

	validate := validator.New()
	userHandler := handlers.NewUserHandler(userService, &utils.Validator{Validator: validate})

	userGroup := app.Group("/users")
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Get("/", middlewares.AuthMiddleware, middlewares.AdminMiddleware, userHandler.GetAllUsers)
	userGroup.Get("/profile", middlewares.AuthMiddleware, userHandler.GetUserProfile)
	userGroup.Put("/profile", middlewares.AuthMiddleware, userHandler.UpdateUserProfile)

	
}