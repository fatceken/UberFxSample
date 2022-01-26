package rediscache

import (
	"github.com/go-redis/redis"
	"time"
)

type Helper struct {
	client *redis.Client
}

func (helper Helper) Add(key string, value interface{}, expiration time.Duration) error {
	err := helper.client.Set(key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (helper Helper) Get(key string) (interface{}, error) {
	val, err := helper.client.Get(key).Result()
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
