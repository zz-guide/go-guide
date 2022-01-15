package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
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

	// 订阅全部消息
	pubsub := rdb.Subscribe(ctx, "name")

	go func() {
		s := <-c
		defer pubsub.Close()
		fmt.Println("退出信号", s)
	}()

	ch := pubsub.Channel()
	for msg := range ch {
		log.Println(msg.Channel, ":", msg.Payload)
	}

	fmt.Println("关闭")
}
