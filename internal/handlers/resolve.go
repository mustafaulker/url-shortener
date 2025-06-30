package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"url-shortener/internal/services"
)

func NewResolveHandler(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")
		u, err := s.ResolveURL(c.Context(), code)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "invalid or expired link"})
		}
		return c.Redirect(u.FullURL, http.StatusFound)
	}
}
