package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result() // 检查是否连接
	if err != nil {
		log.Fatal(err)
	}

	// 连接成功啦
	log.Println(pong)

	if err := rdb.RPush(ctx, "queue", "message").Err(); err != nil {
		panic(err)
	}

	fmt.Println("关闭")
}
