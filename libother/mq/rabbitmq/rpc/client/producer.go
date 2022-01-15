package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_rpc"
	exchangeType = "direct"
	consumerName = "consumer_rpc_client"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/**
使用 MQ 实现 RPC 的意义
通过 MQ 实现 RPC 看起来比客户端和服务器直接通讯要复杂一些，那我们为什么要这样做呢？或者说这样做有什么好处：

将客户端和服务器解耦：客户端只是发布一个请求到 MQ 并消费这个请求的响应。并不关心具体由谁来处理这个请求，MQ 另一端的请求的消费者可以随意替换成任何可以处理请求的服务器，并不影响到客户端。
减轻服务器的压力：传统的 RPC 模式中如果客户端和请求过多，服务器的压力会过大。由 MQ 作为中间件的话，过多的请求而是被 MQ 消化掉，服务器可以控制消费请求的频次，并不会影响到服务器。
服务器的横向扩展更加容易：如果服务器的处理能力不能满足请求的频次，只需要增加服务器来消费 MQ 的消息即可，MQ会帮我们实现消息消费的负载均衡。
可以看出 RabbitMQ 对于 RPC 模式的支持也是比较友好地，
amq.rabbitmq.reply-to, reply_to, correlation_id这些特性都说明了这一点，再加上 spring-rabbit 的实现，可以让我们很简单的使用消息队列模式的 RPC 调用。


1.每一个客户端创建一个响应队列，这个队列应该由客户端来创建且只能由这个客户端使用并在使用完毕后删除，这里可以使用 RabbitMQ 提供的排他队列（Exclusive Queue）
并且要保证队列名唯一，声明队列时名称设为空 RabbitMQ 会生成一个唯一的队列名。


producer生产者相当于是客户端，发起rpc请求
consumer消费者相当于是服务端，监听请求等待处理

producer----绑定随机排他队列---发送消息到consumer的队列
consumer---绑定自己的队列处理完毕消息--相应结果到producer


*/
func main() {
	RpcClient()
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RpcClient() {
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, "")
	if err != nil {
		return
	}

	defer instance.Close()

	q, err := instance.ExclusiveQueueDeclare("")
	if err != nil {
		return
	}

	if err := instance.QueueBind(exchangeName, q.Name, q.Name); err != nil {
		return
	}

	msgs, err := instance.Channel.Consume(
		q.Name,       // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		log.Printf("client consumer error: %s \n", err)
		return
	}

	corrId := randomString(32)

	body := fmt.Sprintf("{'name':'许磊','id':'%d'}", 1)
	err = instance.Channel.Publish(
		exchangeName, // exchange
		"key_rpc",    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(body),
		})
	if err != nil {
		log.Printf("client Publish error: %s \n", err)
		return
	}

	fmt.Println("asdasdasd")
	for d := range msgs {
		log.Printf("接收rpc server 响应: %s \n", string(d.Body))
		if corrId == d.CorrelationId {
			break
		}
	}

	log.Printf("6秒之后关闭连接")
	time.Sleep(time.Second * 6)
	return
}
