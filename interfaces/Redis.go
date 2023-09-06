package interfaces

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}
