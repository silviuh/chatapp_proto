package db

import (
	"github.com/go-redis/redis/v8"
)

// ConnectRedis establishes a connection to Redis.
func ConnectRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr, // e.g. "localhost:6379"
		Password: "",   // no password set
		DB:       0,    // use default DB
	})
}
