package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserCacheInterface interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type UserCache struct {
	Expiration time.Duration
	client     *redis.Client
}

func NewUserCache(redisCli *redis.Client, expiration time.Duration) UserCacheInterface {
	return &UserCache{
		client:     redisCli,
		Expiration: expiration,
	}
}

func (b *UserCache) Get(ctx context.Context, key string) (string, error) {
	value := b.client.Get(ctx, key).Val()

	if value == "" {
		return "", nil
	}

	return value, nil
}

func (b *UserCache) Set(ctx context.Context, key string, value string) error {
	return b.client.Set(ctx, key, value, b.Expiration).Err()
}
