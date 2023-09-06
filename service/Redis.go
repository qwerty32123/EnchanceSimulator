package services

import (
	"EnchanceSimulator/interfaces"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RedisService struct {
	Redis interfaces.RedisClient
}

// ItemStorage is responsible for storing and retrieving items in Redis.
type ItemStorage struct {
	RedisClient *redis.Client
}

// NewItemStorage creates a new ItemStorage with the provided Redis client.
func NewItemStorage(redisClient *redis.Client) *ItemStorage {
	return &ItemStorage{RedisClient: redisClient}
}

func NewRedisService(redisClient interfaces.RedisClient) *RedisService {
	return &RedisService{
		Redis: redisClient,
	}
}

func (is *ItemStorage) StoreItem(ctx context.Context, key string, itemJSON string, expiration time.Duration) bool {
	err := is.RedisClient.Set(ctx, key, itemJSON, expiration).Err()
	if err != nil {
		log.Printf("Error storing item in Redis: %v\n", err)
		return false
	}
	fmt.Println("Item stored in Redis successfully.")
	return true
}
func getItemKey(id int) string {
	return fmt.Sprintf("item:%d", id)
}
