package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"url-shortener/internal/models"
)

type Cache interface {
	Get(ctx context.Context, code string) (*models.ShortURL, error)
	Set(ctx context.Context, url *models.ShortURL, ttl time.Duration) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) Cache {
	return &RedisCache{client: client}
}

// Get retrieves a cached ShortURL from Redis
func (r *RedisCache) Get(ctx context.Context, code string) (*models.ShortURL, error) {
	key := "shorturl:" + code
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// Cache miss
		return nil, nil
	}
	if err != nil {
		// Redis error
		return nil, err
	}

	var url models.ShortURL
	if err := json.Unmarshal([]byte(val), &url); err != nil {
		return nil, err
	}
	return &url, nil
}

// Set stores a ShortURL in Redis with TTL
func (r *RedisCache) Set(ctx context.Context, url *models.ShortURL, ttl time.Duration) error {
	key := "shorturl:" + url.Code
	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}
