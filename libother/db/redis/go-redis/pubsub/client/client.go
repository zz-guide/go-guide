package main

import (
	"context"
	"fmt"
	"go-guide/libother/db/redis/go-redis/common"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	rdb := common.Rdb()
	defer rdb.Close()

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
