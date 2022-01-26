package rediscache

import (
	"github.com/go-redis/redis"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Provide(
		createClient,
		createHelper,
	)
}

func createClient() *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "123456",
			DB:       0, // use default DB
		})

	return client
}

func createHelper(c *redis.Client) *Helper {
	return &Helper{
		client: c,
	}
}
