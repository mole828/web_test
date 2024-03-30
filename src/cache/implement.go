package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type WebCache_RedisImpl struct {
	redisClient *redis.Client
}

func (cache *WebCache_RedisImpl) Inc(key string) (int64, error) {
	cmd := cache.redisClient.Incr(context.Background(), key)
	return cmd.Result()
}

func New_WebCache_RedisImpl(addr string) *WebCache_RedisImpl {
	return &WebCache_RedisImpl{
		redisClient: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}
