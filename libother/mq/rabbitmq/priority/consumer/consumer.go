package main

import (
	"fmt"
	"go-guide/libother/mq/rabbitmq/util"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// exchange,routing key,queue唯一确定一个consumer
var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_pri"
	exchangeType = "topic"
	routingKey   = "key_pri"
	queueName    = "queue_pri"

	consumerName = "consumer_pri"

	instance *util.RabbitmqInstance
)

func main() {
	fmt.Println("消费者启动成功")

	ConsumePush()
	defer instance.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("main主动退出")
}

func ConsumePush() {
	var err error
	instance, err = util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}

	// channel缓冲，慢慢消费
	err = instance.Channel.Qos(1, 0, false)

	// push模式
	messages, err := instance.Channel.Consume(
		queueName,    // queue
		consumerName, // consumer 名称
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	go func() {
		for d := range messages {
			log.Printf("接收信息: %s", string(d.Body))
			d.Ack(false)
			time.Sleep(time.Second * 2)
		}
	}()
}
