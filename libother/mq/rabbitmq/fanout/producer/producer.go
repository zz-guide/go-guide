package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_fanout1"
	exchangeType = "fanout" // 其实就是发布订阅模式
	routingKey   = ""       // 不需要routing key
	queueName    = ""
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	SendMessage()
	SendMessage1()
}

func SendMessage1() {
	var queueName1 = "e1_queue_fanout2"
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName1)
	if err != nil {
		return
	}

	defer instance.Close()
	for i := 0; i < 5; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.Publish(body, exchangeName, routingKey)
	}
}

func SendMessage() {
	var queueName1 = "e1_queue_fanout1"
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName1)
	if err != nil {
		return
	}

	defer instance.Close()
	for i := 0; i < 5; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.Publish(body, exchangeName, routingKey)
	}
}
