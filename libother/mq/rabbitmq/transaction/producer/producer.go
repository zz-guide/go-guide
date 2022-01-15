package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_trx1"
	exchangeType = "direct"
	routingKey   = "key1"
	queueName    = "queue1"
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
	TrxSendMessage()
}

func TrxSendMessage() {
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	defer instance.Close()
	_ = instance.Channel.Tx()

	for i := 0; i < 500; i++ {
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		err = instance.Publish(body, exchangeName, routingKey)
		if err != nil {
			_ = instance.Channel.TxRollback()
		}
	}

	err = instance.Channel.TxCommit()
	if err != nil {
		_ = instance.Channel.TxRollback()
	}
}
