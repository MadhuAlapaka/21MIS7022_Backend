package utils

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

// CacheClient is the Redis client instance
var CacheClient *redis.Client

// InitializeCache sets up the Redis cache client
func InitializeCache(redisAddr string) {
	CacheClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0, // Default DB
	})

	// Test connection
	_, err := CacheClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully!")
}

// SetCache stores data in Redis with an expiration time
func SetCache(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return CacheClient.Set(ctx, key, data, expiration).Err()
}

// GetCache retrieves data from Redis
func GetCache(key string, dest interface{}) error {
	ctx := context.Background()
	data, err := CacheClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

// DeleteCache removes a specific key from Redis
func DeleteCache(key string) error {
	ctx := context.Background()
	return CacheClient.Del(ctx, key).Err()
}

// InvalidateCache clears cache for updated metadata
func InvalidateCache(fileID string) error {
	cacheKey := "file_metadata:" + fileID
	return DeleteCache(cacheKey)
}
