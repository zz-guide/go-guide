package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

var consumerGroup = "bingo-group"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//ReceiveMultiMessage()
	//ReceiveOneMessageUnCommit()
	ReceiveOneMessageManualCommit()
}

func ReceiveMultiMessage() {
	topic := "bingo"
	// 问题：partition不存在一直阻塞
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:19092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 读取超时设置
	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 5))

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println("接收消息:", string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func ReceiveOneMessageUnCommit() {
	topic := "bingo"
	// 问题：partition不存在一直阻塞
	//partition := 0
	ctx := context.Background()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:19092"},
		Topic:   topic,
		GroupID: consumerGroup,
		//Partition: partition,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  2 * time.Second,
		//PartitionWatchInterval: 200 * time.Millisecond,
		ReadLagInterval: 2 * time.Second,
	})

	// 从哪里开始消费，注意：定义了GroupId就不能使用该方法了
	//_ = r.SetOffset(14)

	// 取一条消息，ReadMessage方法底层调用FetchMessage，使用消费者组将自动commit offset
	for {
		log.Printf("开始拉取消息\n")
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Printf("read message error %s \n:", err)
			time.Sleep(time.Second * 2)
			continue
		}

		fmt.Printf("接收消息: %d: %s6 = %s \n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

	fmt.Println("-----消费完毕----")
}

func ReceiveOneMessageManualCommit() {
	topic := "bingo"
	//partition := 0
	ctx := context.Background()

	// Partition和消费者组只能同时设置一个，不能都设置
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          []string{"localhost:19092"},
		Topic:            topic,
		GroupID:          consumerGroup,
		MinBytes:         1,
		MaxBytes:         10e6,
		MaxWait:          100 * time.Millisecond,
		RebalanceTimeout: time.Second,
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
