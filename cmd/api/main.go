package main

import (
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/skyline-digital/api/internal/db"
	"github.com/skyline-digital/api/internal/env"
	"github.com/skyline-digital/api/internal/store"
	"github.com/skyline-digital/api/pkg/api"
)

func main() {
	cfg := api.Config{
		Host: env.GetString("HOST", "localhost"),
		Port: env.GetInt("PORT", 8080),
		DB: api.DbConfig{
			Host:         env.GetString("DB_HOST", "localhost"),
			Port:         env.GetInt("DB_PORT", 5432),
			User:         env.GetString("DB_USER", "admin"),
			Password:     env.GetString("DB_PASSWORD", "adminpassword"),
			Database:     env.GetString("DB_NAME", "skyline"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 10*time.Minute),
		},
	}

	db, err := db.NewPostgres(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Database,
		cfg.DB.MaxOpenConns,
		cfg.DB.MaxIdleConns,
		cfg.DB.MaxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	store := store.NewStorage(db)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
