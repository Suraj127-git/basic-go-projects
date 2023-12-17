package database

import (
	"context"
	"os"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func CreateRedisClient(dbNo int) *redis.Client {
	redis_conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return redis_conn
}
