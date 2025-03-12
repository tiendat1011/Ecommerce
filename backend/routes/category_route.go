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

func CategoryRoutes(app fiber.Router) {
	categoryDAO := daos.NewCategoryDAO()
	categoryService := services.NewCategoryService(categoryDAO)

	validate := validator.New()
	categoryHandler := handlers.NewCategoryHandler(categoryService, &utils.Validator{Validator: validate})

	categoryGroup := app.Group("/category", middlewares.AuthMiddleware)
	categoryGroup.Post("/", categoryHandler.CreateCategory)
	categoryGroup.Put("/:id", categoryHandler.UpdateCategory)
	categoryGroup.Delete("/:id", categoryHandler.DeleteCategory)
}