package rediscache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Helper struct {
	client *redis.Client
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
