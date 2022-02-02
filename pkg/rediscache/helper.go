package rediscache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Helper struct {
	client IRedis
}

func (helper Helper) Add(key string, value interface{}, expiration time.Duration) error {
	err := helper.client.Set(context.Background(), key, value, expiration).Err()

	if err != nil {
		return err
	}
	return nil
}

func (helper Helper) Get(key string) (interface{}, error) {
	val, err := helper.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (helper Helper) Close() error {
	return helper.client.Close()
}

type IRedis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Close() error
}
