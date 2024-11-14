package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/skyline-digital/api/internal/env"
	"github.com/skyline-digital/api/pkg/api"
)

func main() {
	cfg := api.Config{
		Host: env.GetString("HOST", "localhost"),
		Port: env.GetInt("PORT", 8080),
	}

	app := &api.Application{
		Config: cfg,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
