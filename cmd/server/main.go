package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"url-shortener/api"
	"url-shortener/config"
)

func main() {
	cfg, _ := config.Load()
	app := fiber.New()
	app.Use(recover.New(), logger.New())
	api.Setup(app, cfg)
	err := app.Listen(cfg.Port)
	if err != nil {
		return
	}
}
