package db

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

//RDB ...
var (
	rdbClient *redis.Client
	redisOnce sync.Once
)

//InitRedis ...
func InitRedis() *redis.Client {

	redisOnce.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		rdbClient = rdb
	})
	return rdbClient
}
