package database

import (
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(addr, pass string) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: addr, Password: pass})
}
