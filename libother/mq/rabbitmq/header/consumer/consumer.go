package main

import (
	"fmt"
	"go-guide/libother/mq/rabbitmq/util"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

// exchange,routing key,queue唯一确定一个consumer
// header没有routing key
var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_header1"
	exchangeType = "headers"
	routingKey   = ""
	queueName    = "e1_queue_header2"

	consumerName = "consumer2"

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

	// 如果队列的参数中匹配到了消息的header则可以被接收
	//headers := amqp.Table{"mail": "373045134@qq.com"}
	headers := amqp.Table{"author": "许磊1"}
	instance.SetPublishQueueHeaders(headers)

	if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}

	messages, err := instance.Channel.Consume(
		queueName,    // queue
		consumerName, // consumer 名称
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	go func() {
		for d := range messages {
			log.Printf("接收信息: %s", string(d.Body))
		}
	}()
}
