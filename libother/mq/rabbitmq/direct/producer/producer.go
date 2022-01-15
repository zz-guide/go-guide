package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_d1"
	exchangeType = "direct"
	routingKey   = "key_e1_queue_d1"
	queueName    = "e1_queue_1"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	SendMessage()
}

func SendMessage() {
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	defer instance.Close()
	for i := 0; i < 10; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.Publish(body, exchangeName, routingKey)
	}
}
