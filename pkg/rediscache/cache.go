package rediscache

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"os"
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
			Addr:     os.Getenv("MYPREFIX_RedisHost"),
			Password: os.Getenv("MYPREFIX_RedisPassword"),
			DB:       0, // use default DB
		})

	return client
}

func createHelper(c *redis.Client) *Helper {
	return &Helper{
		client: c,
	}
}
