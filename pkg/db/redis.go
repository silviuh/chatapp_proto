package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RDB *redis.Client

// InitializeRedis initializes and returns a Redis client.
func InitializeRedis(redisAddr string) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisAddr, // e.g., "localhost:6379"
		Password: "",        // no password set
		DB:       0,         // use default DB
	})

	// Verify connection
	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully.")
}

func MarkUserOnline(userEmail string) {
	ctx := context.Background()
	err := RDB.Set(ctx, "user_online:"+userEmail, "true", 0).Err()
	if err != nil {
		log.Printf("Failed to mark user as online: %v", err)
	}
}

// IsUserOnline checks if a user is marked as online in Redis by their email.
func IsUserOnline(userEmail string) bool {
	ctx := context.Background()
	result, err := RDB.Get(ctx, "user_online:"+userEmail).Result()
	if err != nil {
		log.Printf("Error checking if user is online: %v", err)
		return false
	}
	return result == "true"
}
