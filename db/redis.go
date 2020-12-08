package db

import (
	"github.com/go-redis/redis/v8"
)

//RDB ...
var RDB *redis.Client

//InitRedis ...
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	RDB = rdb
}
