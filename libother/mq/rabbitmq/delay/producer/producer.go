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
	exchangeName = "exchange_normal_delay"
	exchangeType = "x-delayed-message" // 基于插件实现的延时队列
	routingKey   = "key_normal_delay"
	queueName    = "queue_normal_delay"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/**
1.利用rabbitmq_delayed_message_exchange插件，实现延时队列，不需要死信队列，并且还能实现队列里哪个消息先到期先消费
2.exchange声明type为x-delayed-message
3.message发送在header中声明x-delay
*/

func main() {
	SendMessage()
}

func SendMessage() {
	instance, err := util.NewRabbitmqInstance(uri, "", "", queueName)
	if err != nil {
		return
	}

	if err := instance.DelayExchangeDeclare(exchangeName, exchangeType); err != nil {
		return
	}

	if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}

	defer instance.Close()
	for i := 0; i < 1; i++ {
		time.Sleep(200 * time.Microsecond)
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d', 'delay':'slow'}", num)
		instance.PublishDelay(body, exchangeName, routingKey, amqp.Table{"x-delay": "10000"})
	}

	for i := 0; i < 1; i++ {
		time.Sleep(200 * time.Microsecond)
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d','delay':'fast'}", num)
		instance.PublishDelay(body, exchangeName, routingKey, amqp.Table{"x-delay": "5000"})
	}
}
