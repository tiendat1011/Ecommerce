package routes

import (
	"ecommerce-project/daos"
	"ecommerce-project/handlers"
	"ecommerce-project/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app fiber.Router) {
	productDAO := daos.NewCategoryDAO()
	productService := services.NewProductService(productDAO)
	productHandler := handlers.NewProductHandler(productService)

	fmt.Println(productHandler)
}