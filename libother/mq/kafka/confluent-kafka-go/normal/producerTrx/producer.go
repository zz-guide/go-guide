package main

// 文档链接：https://github.com/confluentinc/confluent-kafka-go
// 所有的配置项：https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md

import (
	"context"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	. "go-guide/libother/mq/kafka/common"
)

/**
事务生产者
*/
func main() {
	TransactionSendMessage()
}

// TransactionSendMessage
// 生产者负责写数据，所以需要事务
// 消费者是读消息，不需要事务
func TransactionSendMessage() {
	topic := "bingo"
	broker := "localhost:19092,localhost:19093,localhost:19094"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":         broker,
		"go.batch.producer":         true, // 生产者是否能批量发送
		"enable.idempotence":        true, // 开启幂等性保护，防止消息重复，开启幂等时，ack必然为all
		"request.required.acks":     -1,   // acks值
		"socket.timeout.ms":         6000, // 网络请求的超时时间
		"message.send.max.retries":  3,    // 消息重试次数
		"reconnect.backoff.max.ms":  3000, // 设置客户端内部重试间隔。
		"retry.backoff.ms":          1000, // 重试时间间隔
		"linger.ms":                 10,   // 发送消息的最大延时时间,ms
		"batch.size":                5,    // 消息最大条数
		"go.delivery.report.fields": "all",
		"transactional.id":          "TR1",  // 事务ID
		"transaction.timeout.ms":    "8000", // 事务超时
	})
	if err != nil {
		log.Printf("生产者创建失败: %s\n", err)
		return
	}

	defer producer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	err = producer.InitTransactions(ctx)
	err = producer.BeginTransaction()

	deliveryChan := make(chan kafka.Event)
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          GetMessage("事务1"),
	}, deliveryChan)

	_, ok := <-deliveryChan
	if !ok {
		log.Printf("发送失败: %s\n", err)
		_ = producer.AbortTransaction(ctx)
		return
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 2},
		Value:          GetMessage("事务2"),
	}, deliveryChan)

	_, ok = <-deliveryChan
	if !ok {
		log.Printf("发送失败: %s\n", err)
		_ = producer.AbortTransaction(ctx)
		return
	}

	err = producer.CommitTransaction(ctx)
	if err != nil {
		_ = producer.AbortTransaction(ctx)
	}

	log.Printf("kafka事务成功")
}
