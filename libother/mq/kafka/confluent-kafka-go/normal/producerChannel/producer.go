package main

// 文档链接：https://github.com/confluentinc/confluent-kafka-go
// 所有的配置项：https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	. "go-guide/libother/mq/kafka/common"
	"log"
)

/**
结论：1.发送都是异步的，有2种写法，基于channel
2.发送只有单条发送，但是可以开启批量发送优化性能，此功能于kafka无关，for循环发即可
3，可配置幂等性和ack等参数保证消息不丢失
*/
func main() {
	SendMessage()
}

func SendMessage() {
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
	})
	if err != nil {
		log.Printf("生产者创建失败: %s\n", err)
		return
	}

	defer producer.Close()

	// 异步发送，可以发送多条
	messages := []*kafka.Message{
		{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          GetMessage("Hello Go"),
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
		},
		{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          GetMessage("Hello111 Go"),
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary1")}},
		},
	}

	done := make(chan bool)
	go func() {
		defer close(done)

		successCnt := 0
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				message := ev
				if message.TopicPartition.Error != nil {
					log.Printf("发送失败: %v\n", message.TopicPartition.Error)
					done <- false
				} else {
					log.Printf("发送消息: %s 到topic %s [%d] at offset %v\n", string(message.Value), *message.TopicPartition.Topic, message.TopicPartition.Partition, message.TopicPartition.Offset)

					successCnt++
					if successCnt == len(messages) {
						done <- true
					}
				}

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	for _, message := range messages {
		producer.ProduceChannel() <- message
	}

	producer.Flush(10 * 1000)
	_ = <-done

	log.Printf("发送完毕: %s\n", err)
}
