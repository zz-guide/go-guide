package main

import (
	"fmt"
	"go-guide/libother/mq/rabbitmq/util"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_d1"
	exchangeType = "direct"
	routingKey   = "key_e1_queue_d1"
	queueName    = "e1_queue_1"

	consumerName = "consumer1"

	instance *util.RabbitmqInstance
)

func main() {
	fmt.Println("消费者启动成功")

	Consume()
	defer instance.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("main主动退出")
}

func Consume() {
	var err error
	instance, err = util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}

	// ack确认机制
	messages, err := instance.Channel.Consume(
		queueName,    // queue
		consumerName, // consumer 名称
		false,        // auto-ack 不自动确认ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	go func() {
		for d := range messages {
			log.Printf("接收信息: %s", string(d.Body))

			d.Ack(false)
		}
	}()
}
