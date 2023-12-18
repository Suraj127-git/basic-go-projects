package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func CreateRedisClient(dbNo int) *redis.Client {
	redis_conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return redis_conn
}
