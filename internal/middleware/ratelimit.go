package middleware

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
	"url-shortener/config"
)

func NewRateLimiter(rdb *redis.Client, cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		key := "ratelimit:" + ip

		ctx := context.Background()
		countStr, err := rdb.Get(ctx, key).Result()

		var count int
		if errors.Is(err, redis.Nil) {
			count = 0
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "rate limiter failed"})
		} else {
			count, _ = strconv.Atoi(countStr)
		}

		if count >= cfg.RateLimitRequests {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, try again later.",
			})
		}

		pipe := rdb.TxPipeline()
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, time.Duration(cfg.RateLimitDurationMinutes)*time.Minute)
		_, _ = pipe.Exec(ctx)

		return c.Next()
	}
}
