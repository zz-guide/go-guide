package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_topic1"
	exchangeType = "topic"
	routingKey   = "key.topic.*" // topic 模式必须用.连接起来，*代替一个词，#代表0个和多个词
	queueName    = "e1_queue_topic1"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	//SendMessage()
	SendMessage1()
}

func SendMessage1() {
	var routingKey1 = "key.topic.one"
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	defer instance.Close()
	for i := 0; i < 2; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.Publish(body, exchangeName, routingKey1)
	}
}

func SendMessage() {
	var routingKey1 = "a.topic.one"
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	defer instance.Close()
	for i := 0; i < 2; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.Publish(body, exchangeName, routingKey1)
	}
}
