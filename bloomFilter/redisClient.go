package bloomFilter

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	c *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	return &RedisClient{
		c: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

// Eval 参数说明：若下标从1开始，keys[1]为 bitMapKeyName；args[1] 为args的长度,其后续为对应的bit位置
func (rdb *RedisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	return rdb.c.Eval(ctx, script, keys, args...).Result()
}
