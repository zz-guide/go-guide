package main

// 文档链接：https://github.com/confluentinc/confluent-kafka-go

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("消费者启动")
	TransactionReceiveMessage()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("消费者关闭")
}

func TransactionReceiveMessage() {
	topics := []string{"bingo"}
	broker := "localhost:19092,localhost:19093,localhost:19094"
	GroupName := "bingo-group1"

	// session.timeout.ms 使用 Kafka 消费分组机制时，消费者超时时间。当 Broker 在该时间内没有收到消费者的心跳时，认为该消费者故障失败，Broker
	// 发起重新 Rebalance 过程。目前该值的配置必须在 Broker 配置group.min.session.timeout.ms=6000和group.max.session.timeout.ms=300000 之间

	//isolation.level 设置为read_committed时候是生产者已提交的数据才能读取到
	//isolation.level 设置为read_uncommitted时候可以读取到未提交的数据
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"api.version.request":       "true",
		"bootstrap.servers":         broker,
		"group.id":                  GroupName,
		"auto.offset.reset":         "latest", // latest, latest, none                                                                                                                                                                                                                                                                                                         latest
		"enable.auto.commit":        false,    // 开启手动提交
		"auto.commit.interval.ms":   5000,     // offset自动提交间隔
		"heartbeat.interval.ms":     3000,     // 心跳时间
		"session.timeout.ms":        25000,    // 心跳超时时间，超过该时间认为消费端不可用
		"max.poll.interval.ms":      30000,    // 最大pull时间
		"fetch.max.bytes":           1024000,
		"max.partition.fetch.bytes": 256000,
		"reconnect.backoff.max.ms":  1000,
		"isolation.level":           "read_committed", // 支持两种事务隔离级别：read_uncommitted,read_committed
	})

	if err != nil {
		log.Printf("消费者创建失败: %s\n", err)
		return
	}

	// 获取订阅TopicPartition的最新消费位置
	offsets, err := consumer.Position([]kafka.TopicPartition{
		{Topic: &topics[0], Partition: 0},
		{Topic: &topics[0], Partition: 2},
	})

	//consumer.IncrementalAssign()

	// 开始消费
	err = consumer.Assign(offsets)

	// 直接跳转到指定offset开始消费
	//consumer.Seek()

	/*err = consumer.SubscribeTopics(topics, func(c *kafka.Consumer, e kafka.Event) error {
		log.Println("rebalance中:", e)
		return nil
	})
	*/

	//consumer.Subscribe(topics[0], nil)

	if err != nil {
		log.Printf("消费者订阅主题失败: %s\n", err)
		return
	}

	go func() {
		defer consumer.Close()
		for {
			log.Println("新的一轮")

			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("没有获取到消息，重试: %v (%v)\n", err, msg)
				continue
			}

			log.Printf("接收消息 %s %s %v\n", msg.TopicPartition, string(msg.Value), msg.Headers)

			_, _ = consumer.CommitMessage(msg)

			// 这几种都可以实现提交
			//consumer.CommitOffsets()
			//consumer.CommitMessage()
			//consumer.Committed()
			//consumer.Commit()
		}
	}()
}
