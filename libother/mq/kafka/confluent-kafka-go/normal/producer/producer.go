package main

// 文档链接：https://github.com/confluentinc/confluent-kafka-go

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	. "go-guide/libother/mq/kafka/common"
)

func main() {
	SendMessage()
}

func SendMessage() {
	topic := "bingo"
	broker := "localhost:19092,localhost:19093,localhost:19094"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"enable.idempotence": true, // 开启幂等性保护，防止消息重复
	})
	if err != nil {
		log.Printf("生产者创建失败: %s\n", err)
		return
	}

	defer producer.Close()
	deliveryChan := make(chan kafka.Event)

	// 异步发送,只能发送一条
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          GetMessage("Hello Go"),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)

	e, ok := <-deliveryChan
	close(deliveryChan)
	if !ok {
		log.Printf("发送失败: %s\n", err)
		return
	}

	message := e.(*kafka.Message)
	log.Printf("发送消息到topic %s [%d] at offset %v\n", *message.TopicPartition.Topic, message.TopicPartition.Partition, message.TopicPartition.Offset)
}
