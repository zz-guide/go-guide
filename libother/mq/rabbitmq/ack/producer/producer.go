package main

import (
	"fmt"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_d1"
	exchangeType = "direct"
	routingKey   = "key_e1_queue_d2"
	queueName    = "e1_queue_2"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/**
为了保证发送方消息可靠发送到rabbitmq，官方提供2种机制
1.事务机制
性能慢，因为网络交互会变多

2.confirm模式

*/
func main() {
	SendMessage()
	//time.Sleep(10 * time.Second)
}

func SendMessage() {
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	defer instance.Close()

	// 生产者ack模式
	_ = instance.Channel.Confirm(false)

	confirms := instance.Channel.NotifyPublish(make(chan amqp.Confirmation, 10))
	defer confirmOne(confirms) // 处理方法

	for i := 0; i < 500; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.NotifyPublish(body, exchangeName, routingKey)
	}
}

func confirmOne(confirms <-chan amqp.Confirmation) {
	if confirmed := <-confirms; confirmed.Ack {
		fmt.Printf("ack true: %d\n", confirmed.DeliveryTag)
	} else {
		fmt.Printf("ack false: %d\n", confirmed.DeliveryTag)
	}
}
