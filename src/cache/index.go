package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var redisMaster *redis.Client

func init() {
	redisMaster = redis.NewClient(&redis.Options{
		Addr: "redis-master-service",
	})
}

func Inc() (int64, error) {
	int_cmd := redisMaster.Incr(context.Background(), "count")
	return int_cmd.Result()
}
