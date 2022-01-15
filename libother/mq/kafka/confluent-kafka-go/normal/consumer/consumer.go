package main

// 文档链接：https://github.com/confluentinc/confluent-kafka-go

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("消费者启动")
	ReceiveMessage1()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("消费者关闭")
}

func ReceiveMessage() {
	topics := []string{"bingo"}
	broker := "localhost:19092,localhost:19093,localhost:19094"
	GroupName := "bingo-group1"

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"api.version.request": "true",
		"bootstrap.servers":   broker,
		"group.id":            GroupName,
	})

	if err != nil {
		log.Printf("消费者创建失败: %s\n", err)
		return
	}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Printf("消费者订阅主题失败: %s\n", err)
		return
	}

	go func() {
		run := true

		for run {
			log.Println("新的一轮")
			select {
			default:
				ev := consumer.Poll(2000)
				if ev == nil {
					continue
				}

				switch e := ev.(type) {
				case *kafka.Message:
					log.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
					if e.Headers != nil {
						log.Printf("%% Headers: %v\n", e.Headers)
					}
				case kafka.Error:
					fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
					if e.Code() == kafka.ErrAllBrokersDown {
						run = false
					}
				default:
					fmt.Printf("Ignored %v\n", e)
				}
			}
		}

		log.Printf("Closing consumer\n")
		consumer.Close()
	}()
}

func ReceiveMessage1() {
	topics := []string{"bingo"}
	broker := "localhost:19092,localhost:19093,localhost:19094"
	GroupName := "bingo-group1"

	// session.timeout.ms 使用 Kafka 消费分组机制时，消费者超时时间。当 Broker 在该时间内没有收到消费者的心跳时，认为该消费者故障失败，Broker
	// 发起重新 Rebalance 过程。目前该值的配置必须在 Broker 配置group.min.session.timeout.ms=6000和group.max.session.timeout.ms=300000 之间

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"api.version.request":       "true",
		"bootstrap.servers":         broker,
		"group.id":                  GroupName,
		"auto.offset.reset":         "earliest", // earliest, latest, none                                                                                                                                                                                                                                                                                                         latest
		"enable.auto.commit":        false,      // 开启手动提交
		"auto.commit.interval.ms":   5000,       // offset自动提交间隔
		"heartbeat.interval.ms":     3000,       // 心跳时间
		"session.timeout.ms":        25000,      // 心跳超时时间，超过该时间认为消费端不可用
		"max.poll.interval.ms":      30000,      // 最大pull时间
		"fetch.max.bytes":           1024000,
		"max.partition.fetch.bytes": 256000,
		"reconnect.backoff.max.ms":  1000,
	})

	if err != nil {
		log.Printf("消费者创建失败: %s\n", err)
		return
	}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Printf("消费者订阅主题失败: %s\n", err)
		return
	}

	go func() {
		run := true
		for run {
			log.Println("新的一轮")

			// 底层还是调用的Poll,首次会很慢，
			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("没有获取到消息，重试: %v (%v)\n", err, msg)
				continue
			}

			log.Printf("接收消息 %s %s %v\n", msg.TopicPartition, string(msg.Value), msg.Headers)

			_, _ = consumer.CommitMessage(msg)
		}

		log.Printf("Closing consumer\n")
		consumer.Close()
	}()
}
