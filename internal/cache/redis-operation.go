package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func Set(ctx context.Context, key string, value interface{}, expiration Expiration) error {

	// Use default expiration if the provided expiration is 0
	if expiration == 0 {
		expiration = Expiration(ONE_DAY)
	}

	// Marshal the value into JSON
	valueJSON, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	// Set value in Redis with expiration
	err = RedisClient.Set(ctx, strings.ToUpper(key), valueJSON, time.Duration(expiration)).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache in Redis: %w", err)
	}
	return nil
}

func Get(ctx context.Context, key string, dest interface{}) error {
	// Get value from Redis
	result, err := RedisClient.Get(ctx, strings.ToUpper(key)).Result()

	if err == redis.Nil {
		return fmt.Errorf("cache miss: key not found")
	} else if err != nil {
		return fmt.Errorf("failed to get cache from Redis: %w", err)
	}

	// Unmarshal the result into the destination object
	err = json.Unmarshal([]byte(result), dest)
	if err != nil {
		return fmt.Errorf("failed to unmarshal cache data: %w", err)
	}

	return nil
}
