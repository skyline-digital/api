package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/skyline-digital/api/internal/env"
	"github.com/skyline-digital/api/internal/store"
	"github.com/skyline-digital/api/pkg/api"
)

func main() {
	cfg := api.Config{
		Host: env.GetString("HOST", "localhost"),
		Port: env.GetInt("PORT", 8080),
	}

	store := store.NewStorage(nil)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
