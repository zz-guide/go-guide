package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

var consumerGroup = "bingo-group2"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ReceiveOneMessageManualCommit()
}

func ReceiveOneMessageManualCommit() {
	topic := "bingo"
	//partition := 0
	ctx := context.Background()

	// Partition和消费者组只能同时设置一个，不能都设置
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:       []string{"localhost:19092"},
		Topic:         topic,
		GroupID:       consumerGroup,
		MinBytes:      1,
		MaxBytes:      10e6,
		MaxWait:       100 * time.Millisecond,
		RetentionTime: 1,
		//PartitionWatchInterval: 200 * time.Millisecond,
		//ReadLagInterval: 20 * time.Second,
	})

	// FetchMessage取一条消息，不会自动commit offset
	for {
		log.Printf("开始拉取消息\n")
		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Printf("fetch message error %s \n:", err)
			time.Sleep(time.Second * 2)
			continue
		}

		fmt.Printf("接收消息: %d: %s6 = %s \n", m.Offset, string(m.Key), string(m.Value))

		if err := r.CommitMessages(ctx, m); err != nil {
			log.Printf("commit messages失败: %s \n", err)
		}

		time.Sleep(time.Second * 3) // 3秒处理一次
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

	fmt.Println("-----消费完毕----")
}
