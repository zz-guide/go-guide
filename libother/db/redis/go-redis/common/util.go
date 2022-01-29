package common

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func Rdb() *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:         "127.0.0.1:6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     "123456",
			PoolSize:     10,
			DB:           0,
		},
	)

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Println("Ping出错:", err)
		return nil
	}

	return rdb
}
