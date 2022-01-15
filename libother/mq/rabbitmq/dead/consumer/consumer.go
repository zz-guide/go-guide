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

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_normal"
	exchangeType = "direct"
	routingKey   = "key_normal"
	queueName    = "queue_normal"

	consumerName = "consumer_normal1"

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
	instance, err = util.NewRabbitmqInstance(uri, exchangeName, exchangeType, "")
	if err != nil {
		return
	}

	/*if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}*/

	// 确保rabbitMQ一个一个发送消息
	err = instance.Channel.Qos(1, 0, false)

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
			// 模拟处理时间过长。迟迟不ACK
			log.Printf("正在处理。。。。。。16秒\n")
			time.Sleep(time.Second * 16)
			d.Ack(false)
			log.Printf("已ack,处理完毕\n")
		}
	}()
}
