package redis

import (
	"github.com/Snorkin/auth_service/config"
	"github.com/go-redis/redis/v8"
	"time"
)

func CreateRedisClient(cfg *config.Config) *redis.Client {
	addr := cfg.Redis.Address

	if addr == "" {
		addr = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB, // 0 for default
	})
	return client
}
