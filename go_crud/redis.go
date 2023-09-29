package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
}

// GetUser retrieves a user by ID from Redis
func GetUser(id string) (*User, error) {
	ctx := context.Background()
	val, err := client.HGetAll(ctx, fmt.Sprintf("user:%s", id)).Result()
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:    id,
		Name:  val["name"],
		Email: val["email"],
	}

	return user, nil
}

// Implement other Redis operations (Create, Update, Delete)
