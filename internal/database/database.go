package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

var (
	ctx          = context.TODO()
	redisConnMap = make(map[int]*redis.Client)
	redisConnMu  sync.Mutex
)

func CreateRedisClient(dbNo int) *redis.Client {
	redisConnMu.Lock()
	defer redisConnMu.Unlock()

	if client, ok := redisConnMap[dbNo]; ok {
		return client
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})

	if _, err := redisConn.Ping().Result(); err != nil {
		return nil
	}

	redisConnMap[dbNo] = redisConn
	return redisConn
}
