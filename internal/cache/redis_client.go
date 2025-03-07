package cache

import (
	"context"
	"os"

	"example.com/task-managment/internal/utils"
	"github.com/fatih/color"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	redisPath := os.Getenv("REDIS_PATH")
	redisPort := os.Getenv("REDIS_PORT")
	redisAddress := redisPath + ":" + redisPort

	if (redisPath == "") && (redisPort == "") {
		color.Yellow("REDIS_ADDRESS not found in .env, using default redis address")
		redisAddress = "localhost:6379"
	}

	utils.LogSuccess("Redis Address is: ", redisAddress)

	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisAddress,
		DB:   0,
	})

	// Test Redis connection
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		color.Red("Failed to connect to Redis!")
		return err
	}

	utils.LogSuccess("Redis connected successfully!")

	return nil
}
