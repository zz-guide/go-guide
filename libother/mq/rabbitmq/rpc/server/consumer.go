package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_rpc"
	exchangeType = "direct"
	routingKey   = "key_rpc"
	queueName    = "queue_rpc"

	consumerName = "consumer_rpc_server"

	instance *util.RabbitmqInstance
)

func main() {
	fmt.Println("消费者rpc服务端启动成功")

	Consume()
	defer instance.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("main主动退出")
}

func Consume() {
	var err error
	instance, err = util.NewRabbitmqInstance(uri, exchangeName, exchangeType, "")
	if err != nil {
		return
	}

	if err := instance.QueueDeclare(queueName); err != nil {
		return
	}

	if err := instance.QueueBind(exchangeName, queueName, routingKey); err != nil {
		return
	}

	/*err = instance.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)*/

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

			log.Printf("接收信息: %s %s \n", string(d.Body), d.ReplyTo)

			err = instance.Channel.Publish(
				exchangeName, // exchange
				d.ReplyTo,    // routing key
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte("---响应---"),
				})

			if err != nil {
				log.Printf("server Publish error: %s \n", err)
				return
			}

			// 这里需要注意。如果一个消息设置了手动确认，就必须应答或者拒绝，否则会一直阻塞
			// 没有设置手动确认然后调用该方法也会阻塞
			d.Ack(false)
		}
	}()

}
