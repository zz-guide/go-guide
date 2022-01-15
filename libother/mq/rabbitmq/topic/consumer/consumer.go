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
	exchangeName = "exchange_topic1"
	exchangeType = "topic"
	routingKey   = "key.topic.*" // topic 模式必须用.连接起来，*代替一个词，#代表0个和多个词
	queueName    = "e1_queue_topic1"

	consumerName = "consumer1"

	instance *util.RabbitmqInstance
)

func main() {
	fmt.Println("消费者启动成功")

	//ConsumePush()
	ConsumePull()
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
		}
	}()
}

func ConsumePull() {
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

	// pull 模式,只能一条一条获取，不能批量拉取
	go func() {
		for {
			log.Printf("新的一轮拉取消息----\n")
			// pull模式,官方建议使用Consume方法
			messages, ok, err := instance.Channel.Get(queueName, false)
			if err != nil {
				log.Printf("pull消息出错: %s \n", err)
				break
			}

			if ok {
				log.Printf("接收信息: %s", string(messages.Body))
				messages.Ack(false)
				continue
			} else {
				time.Sleep(time.Second * 3)
			}
		}

	}()
}
