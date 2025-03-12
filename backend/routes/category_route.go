package routes

import (
	"ecommerce-project/daos"
	"ecommerce-project/handlers"
	"ecommerce-project/middlewares"
	"ecommerce-project/services"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app fiber.Router) {
	categoryDAO := daos.NewCategoryDAO()
	categoryService := services.NewCategoryService(categoryDAO)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	categoryGroup := app.Group("/category", middlewares.AuthMiddleware)
	categoryGroup.Post("/", categoryHandler.CreateCategory)
}
