package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"ecommerce-project/config"
	"ecommerce-project/databases"
	"ecommerce-project/middlewares"
	"ecommerce-project/routes"
)

func main() {
	// Init app
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connection string mongodb
	connectionUri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.Cfg.DbUser,
		config.Cfg.DbPass,
		config.Cfg.DbHost,
		config.Cfg.DbPort,
	)

	databases.InitMongoDB(connectionUri)
	databases.InitValkey()

	app := fiber.New()

	// Middleware
	app.Use(middlewares.SetupLogger())

	// Init Route
	api := app.Group("/api/v1")
	routes.UserRoutes(api)
	routes.AuthRoute(api)
	routes.CategoryRoutes(api)

	// Start server
	serverAddr := fmt.Sprintf("%s:%s", "", config.Cfg.ServerPort)
	log.Printf(serverAddr)
	log.Printf("Server running on port %s", config.Cfg.ServerPort)
	log.Fatal(app.Listen(serverAddr))
}
