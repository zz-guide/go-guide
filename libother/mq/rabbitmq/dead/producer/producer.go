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
	exchangeName = "exchange_normal"
	exchangeType = "direct"
	routingKey   = "key_normal"
	queueName    = "queue_normal"

	deadExchangeName = "exchange_dlx"
	deadQueueName    = "queue_dlx"
	deadRoutingKey   = "key_dlx"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/**
1.队列设置 x-message-ttl,单位:毫秒
如果设置了队列的TTL属性，那么一旦消息过期，就会被队列丢弃。过期的 message 总是会出现在 queue 的头部。

2.消息自身设置Expiration参数
假设轮到当前消息要被消费者处理，会先此时才会判定是否过期，过期要么丢弃或者进入死信队列，没过期才会被接收处理
如果此时消费者比较繁忙，即使队列中的消息已经过期了，也不会被立马丢弃，带来的问题就是不够及时的相应业务


另外，还需要注意的一点是，如果不设置TTL，表示消息永远不会过期，如果将TTL设置为0，则表示除非此时可以直接投递该消息到消费者，否则该消息将会被丢弃。
4.还可以使用策略为全部队列设置ttl, rabbitmqctl set_policy TTL “*”。 “{” “的消息的TTL” “：60000}” - -apply -to队列
5.如果消息和队列同时设置ttl，以最低的为准

1.通过一个正常队列+死信队列可以实现等同于专门的延时队列的效果，但是会有问题，消息不能按照时间正常过期
问题：ttl长的如果在前边，ttl短的不会主动判定为过期
2.正常队列负责接收消息，但不设置消费者消费，同时message设置好Expiration，等到了过期时间自动进入死信队列


*/

func main() {
	SendMessage()
}

func SendMessage() {
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, "")
	if err != nil {
		return
	}

	if err := instance.ExchangeDeclare(deadExchangeName, exchangeType); err != nil {
		return
	}

	args := amqp.Table{}
	// 生产者设置队列消息存活时间,毫秒
	// args["x-message-ttl"] = 6000
	//设置死信交换器
	args["x-dead-letter-exchange"] = deadExchangeName
	//设置死信交换器Key,不设置将使用原先队列的routing key
	args["x-dead-letter-routing-key"] = deadRoutingKey

	// 正常队列和死信交换器绑定
	if err := instance.QueueDeclareBindDeadExchange(queueName, args); err != nil {
		return
	}

	if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}

	if err := instance.QueueDeclare(deadQueueName); err != nil {
		return
	}

	if err := instance.QueueBind(deadExchangeName, deadQueueName, deadRoutingKey); err != nil {
		return
	}

	defer instance.Close()
	for i := 0; i < 1; i++ {
		time.Sleep(500 * time.Microsecond)
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.PublishExpiration(body, exchangeName, routingKey, "6000")
	}

	for i := 0; i < 1; i++ {
		time.Sleep(500 * time.Microsecond)
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.PublishExpiration(body, exchangeName, routingKey, "15000")
	}

	for i := 0; i < 1; i++ {
		time.Sleep(500 * time.Microsecond)
		num := rand.Intn(10000)
		body := fmt.Sprintf("{'name':'许磊','id':'%d'}", num)
		instance.PublishExpiration(body, exchangeName, routingKey, "4000")
	}
}
