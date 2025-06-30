package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	DBUrl                    string
	Port                     string
	BaseURL                  string
	RedisAddr                string
	RedisPassword            string
	RateLimitRequests        int
	RateLimitDurationMinutes int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	rateLimitRequests, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_REQUESTS"))
	rateLimitMinutes, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_MINUTES"))

	return &Config{
		DBUrl:                    os.Getenv("DATABASE_URL"),
		Port:                     os.Getenv("PORT"),
		BaseURL:                  os.Getenv("BASE_URL"),
		RedisAddr:                os.Getenv("REDIS_ADDR"),
		RedisPassword:            os.Getenv("REDIS_PASSWORD"),
		RateLimitRequests:        rateLimitRequests,
		RateLimitDurationMinutes: rateLimitMinutes,
	}, nil
}
