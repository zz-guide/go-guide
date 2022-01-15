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
	exchangeName = "exchange_normal_delay"
	exchangeType = "x-delayed-message"
	routingKey   = "key_normal_delay"
	queueName    = "queue_normal_delay"

	consumerName = "consumer_normal_delay1"

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
	// 队列在生产者方已经定义好了
	instance, err = util.NewRabbitmqInstance(uri, "", "", queueName)
	if err != nil {
		return
	}

	if err := instance.DelayExchangeDeclare(exchangeName, exchangeType); err != nil {
		return
	}

	if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}

	// 确保rabbitMQ一个一个发送消息
	//err = instance.Channel.Qos(1, 0, false)

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
