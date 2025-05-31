package config

import (
	"context"
	"go-starter-kit/pkg/logger"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
		Username: os.Getenv("REDIS_USER"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		logger.LogError("Failed to connect to Redis", err)
	} else {
		logger.LogInfo("Connected to Redis successfully")
	}

	// if err := RedisClient.FlushAll(ctx).Err(); err != nil {
	// 	logger.LogError("Failed to flush Redis database", err)
	// } else {
	// 	logger.LogInfo("Redis database flushed successfully")
	// }
}

func CloseRedis() {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RedisClient.Close(); err != nil {
		logger.LogError("Failed to close Redis connection", err)
	} else {
		logger.LogInfo("Redis connection closed successfully")
	}
}
