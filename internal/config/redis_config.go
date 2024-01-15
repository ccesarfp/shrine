package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

// NewRedisClient creates a new instance of client Redis
func NewRedisClient() (*redis.Client, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	db := 0

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()

	return client, err
}
