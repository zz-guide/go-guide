package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-guide/libother/mq/rabbitmq/util"
)

var (
	uri          = "amqp://xulei:123456@localhost:5672/"
	exchangeName = "exchange_pri"
	exchangeType = "topic"
	routingKey   = "key_pri"
	queueName    = "queue_pri"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/**
1.只有当消费者不足，不能及时进行消费的情况下，优先级队列才会生效

验证方式：触发为及时消费场景，常用场景与Qos结合使用
1、可先发送消息，再进行消费
2、开启手动应答、设置Qos。若为1，在一个消费者存在的情况下，除第一个消息外均按优先级进行消费(第一个消息被及时消费掉了)
3、可在方式二的基础上不断增加消费者，也符合优先调用规则


1.首先队列创建设置好最大优先级，例如10
2.消息也要设置优先级，数字越大优先级越高，0-9
3.队列、消息上均要设置优先级才可生效，以较小值为准
参考链接：https://www.iteye.com/blog/sheungxin-2344874
*/
func main() {
	SendMessage()
}

func SendMessage() {
	instance, err := util.NewRabbitmqInstance(uri, exchangeName, exchangeType, queueName)
	if err != nil {
		return
	}

	defer instance.Close()

	// 测试优先级队列
	for i := 0; i < 10; i++ {
		num := rand.Intn(10000)
		if i%2 == 0 {
			body := fmt.Sprintf("{'id':'%d','pri':'2'}", num)
			instance.PublishPriority(body, exchangeName, routingKey, 2)
		} else {
			body := fmt.Sprintf("{'id':'%d','pri':'6'}", num)
			instance.PublishPriority(body, exchangeName, routingKey, 6)
		}

	}
}
