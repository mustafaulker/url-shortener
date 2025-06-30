package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"url-shortener/internal/services"
)

type ShortenRequest struct {
	URL    string `json:"url"`
	Expiry int    `json:"expiry,omitempty"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
	Expiry   int    `json:"expiry_hours"`
}

func NewShortenHandler(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req ShortenRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}
		short, err := s.CreateShortURL(c.Context(), req.URL, req.Expiry)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(ShortenResponse{
			ShortURL: s.GetBaseURL() + "/" + short.Code,
			Expiry:   req.Expiry,
		})
	}
}
