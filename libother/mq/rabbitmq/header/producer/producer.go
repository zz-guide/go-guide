package main

import (
	"fmt"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go-guide/libother/mq/rabbitmq/util"
)

// 参考链接：// https://www.jianshu.com/p/469f4608ce5d

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_header1"
	exchangeType = "headers" // header模式
	routingKey   = ""        // 不需要routing key
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
	var queueName1 = "e1_queue_header2"
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
	var queueName1 = "e1_queue_header1"
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName1)
	if err != nil {
		return
	}

	defer instance.Close()

	// x-match: any all
	headers := amqp.Table{"x-match": "any", "mail": "373045134@qq.com", "author": "许磊"} // 头部信息 any:匹配一个即可 all:全部匹配
	instance.SetPublishHeaders(headers)

	for i := 0; i < 5; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.Publish(body, exchangeName, routingKey)
	}
}
