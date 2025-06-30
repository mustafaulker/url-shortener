package api

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/internal/handlers"
	"url-shortener/internal/repositories"
	"url-shortener/internal/services"
)

func Setup(app *fiber.App, cfg *config.Config) {
	db, err := database.Connect(cfg.DBUrl)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	rdb := database.ConnectRedis(cfg.RedisAddr, cfg.RedisPassword)
	store := repositories.NewPostgresStore(db)
	cache := repositories.NewRedisCache(rdb)
	svc := services.NewService(store, cache, cfg.BaseURL)

	api := app.Group("/api")
	api.Post("/shorten", handlers.NewShortenHandler(svc))

	app.Get("/:code", handlers.NewResolveHandler(svc))
}
