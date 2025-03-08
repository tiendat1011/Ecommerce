package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"ecommerce-project/config"
	"ecommerce-project/databases"
	"ecommerce-project/routes"
)

func main() {
	// Init app
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connection string mongodb
	connectionUri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
	)
	databases.Init(connectionUri, cfg)
	app := fiber.New()

	// Middleware

	// Init Route
	api := app.Group("/api/v1")
	routes.UserRoutes(api)

	// Start server
	serverAddr := fmt.Sprintf("%s:%s", "", cfg.ServerPort)
	log.Printf(serverAddr)
	log.Printf("Server running on port %s", cfg.ServerPort)
	log.Fatal(app.Listen(serverAddr))
}
