package main

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result() // 检查是否连接
	if err != nil {
		log.Fatal(err)
	}

	log.Println(pong)

	for i := 1; i < 1001; i++ {
		err = rdb.Publish(ctx, "name", "hello"+strconv.Itoa(i)).Err()
		if err != nil {
			panic(err)
		}
	}
}
