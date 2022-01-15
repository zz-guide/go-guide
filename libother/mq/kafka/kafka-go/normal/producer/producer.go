package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
	. "go-guide/libother/mq/kafka/common"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//SendMessage()
	//RandomPartitionSendMessage()
	MultiTopicSendMessage()
}

func SendMessage() {
	topic := "bingo"
	// 问题：partition不存在一直阻塞
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:19092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 写入超时设置
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 5))

	// 1.可以支持批量发送消息，返回写入消息的字节数，要么成功，要么失败，自带确认
	bytes, err := conn.WriteMessages(
		kafka.Message{Value: GetMessage("hello")},
		kafka.Message{Value: GetMessage("hello1")},
	)

	if err != nil {
		log.Printf("发送失败: %s \n", err)
		return
	}

	log.Printf("发送成功，字节数: %d \n", bytes)

	defer conn.Close()
}

func newKafkaWriter(kafkaURL []string, topic string) *kafka.Writer {
	// Balancer: 寻找partition策略：RoundRobin轮询，Hash哈希，LeastBytes最少使用，CRC32Balancer一致性hash，Murmur2Balancer一致性hash
	// RequireAll: 三种，默认是0，直接返回
	//  RequireNone (0)  fire-and-forget, do not wait for acknowledgements from the
	//  RequireOne  (1)  wait for the leader to acknowledge the writes
	//  RequireAll  (-1) wait for the full ISR to acknowledge the writes

	// BatchBytes 每条消息的最大字节数
	// BatchSize 一次请求最大发送消息数
	return &kafka.Writer{
		Addr:         kafka.TCP(kafkaURL...),
		Topic:        topic,
		Balancer:     &kafka.RoundRobin{},
		WriteTimeout: time.Second * 5,
		Async:        true,
		BatchSize:    10,
		RequiredAcks: kafka.RequireAll,
	}
}

func RandomPartitionSendMessage() {
	topic := "bingo"
	ctx := context.Background()
	//单机
	//conn := newKafkaWriter([]string{"localhost:19092"}, topic)
	// 集群
	conn := newKafkaWriter([]string{"localhost:19092", "localhost:19093", "localhost:19094"}, topic)

	// 1.可以支持批量发送消息，要么成功，要么失败，自带确认
	// 2.Key用来计算发送到哪个partition，虽然有策略，也可能每次都不用
	err := conn.WriteMessages(
		ctx,
		kafka.Message{Value: GetMessage("random partition1"), Key: []byte("key1")},
		kafka.Message{Value: GetMessage("random partition2"), Key: []byte("key2")},
		kafka.Message{Value: GetMessage("random partition3"), Key: []byte("key3")},
	)

	if err != nil {
		log.Printf("发送失败: %s \n", err)
		return
	}

	log.Printf("发送成功\n")

	defer conn.Close()
}

func MultiTopicSendMessage() {
	topic := "bingo"
	topic1 := "gauss"
	ctx := context.Background()
	//单机
	//conn := newKafkaWriter([]string{"localhost:19092"}, topic)
	// 集群
	conn := newKafkaWriter([]string{"localhost:19092", "localhost:19093", "localhost:19094"}, "")

	// 一次操作可以发送消息到不同的topic
	err := conn.WriteMessages(
		ctx,
		kafka.Message{Value: GetMessage("写入bingo"), Key: []byte("key1"), Topic: topic},
		kafka.Message{Value: GetMessage("写入gauss"), Key: []byte("key2"), Topic: topic1},
	)

	if err != nil {
		log.Printf("发送失败: %s \n", err)
		return
	}

	log.Printf("发送成功\n")

	defer conn.Close()
}
