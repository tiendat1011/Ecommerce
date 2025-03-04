package main

import (
	"fmt"
	"log"

	"ecommerce-project/config"
	db "ecommerce-project/databases"
)

func main() {
	// Loading environment from .env
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
	db.Init(connectionUri)

	return
}
