package util

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitmqInstance struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   map[string]*amqp.Queue

	Headers      amqp.Table
	QueueHeaders amqp.Table
}

func NewRabbitmqInstance(addr, exchangeName, exchangeType, queueName string) (*RabbitmqInstance, error) {
	// 1.addr "amqp://xulei:123456@localhost:5672/"
	// 2.生产者不需要指定队列和routing key的绑定关系
	// 3.exchange为空时发送到默认队列，此时默认为direct,且可以不指定队列
	instance := &RabbitmqInstance{
		Queue: make(map[string]*amqp.Queue),
	}

	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Printf("获取Connection失败: %s \n", err)
		return nil, err
	}

	instance.Conn = conn

	channel, err := conn.Channel()
	if err != nil {
		log.Printf("获取channel失败: %s \n", err)
		return nil, err
	}

	instance.Channel = channel

	if exchangeName != "" {
		if err := instance.ExchangeDeclare(exchangeName, exchangeType); err != nil {
			return nil, err
		}
	}

	if queueName != "" {
		if err := instance.QueueDeclare(queueName); err != nil {
			return nil, err
		}
	}

	return instance, nil
}

func (r *RabbitmqInstance) ExchangeDeclare(name, kind string) error {
	// "Exchange type - direct|fanout|topic|x-custom"
	// 1.direct, 此模式下routing key名称默认和queue name一致，如果不强制的话，默认exchange name 为 ""
	// 2.topic,  将路由键和某模式进行匹配。此时队列需要绑定要一个模式上。符号“#”匹配一个或多个词，符号“*”只能匹配一个词。
	// 3.fanout, 不处理路由键。你只需要简单的将队列绑定到交换机上。一个发送到该类型交换机的消息都会被广播到与该交换机绑定的所有队列上。
	// 4.header, 不处理路由键，而是根据发送的消息内容中的headers属性进行匹配。在绑定Queue与Exchange时指定一组键值对；当消息发送到RabbitMQ时会取到该消息的headers与Exchange绑定时指定的键值对进行匹配；如果完全匹配则消息会路由到该队列，否则不会路由到该队列。headers属性是一个键值对，可以是Hashtable，键值对的值可以是任何类型。而fanout，direct，topic 的路由键都需要要字符串形式的

	durable := true     // 持久化
	autoDelete := false // 不自动删除
	internal := false   // 不是内置，设置是否是RabbitMQ内部使用，默认false。如果设置为 true ，则表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式。
	noWait := false     // 不需要等待应答
	if err := r.Channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, nil); err != nil {
		log.Printf("Exchange %s 创建失败: %s \n", name, err)
		return err
	}

	return nil
}

func (r *RabbitmqInstance) DelayExchangeDeclare(name, kind string) error {
	durable := true     // 持久化
	autoDelete := false // 不自动删除
	internal := false   // 不是内置，设置是否是RabbitMQ内部使用，默认false。如果设置为 true ，则表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式。
	noWait := false     // 不需要等待应答
	args := amqp.Table{"x-delayed-type": "direct"}
	if err := r.Channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args); err != nil {
		log.Printf("Exchange %s 创建失败: %s \n", name, err)
		return err
	}

	return nil
}

func (r *RabbitmqInstance) QueueDeclare(name string) error {
	if _, ok := r.Queue[name]; ok {
		log.Printf("Queue %s 已存在 \n", name)
		return errors.New("Queue " + name + "已存在")
	}

	durable := true     // 持久化
	autoDelete := false // 不自动删除
	exclusive := false  // 排他，其他用户不可见此队列，断开连接自动删除
	noWait := false     // 不需要等待应答，默认是需要等待

	args := amqp.Table{}
	// args["x-max-length"] = 1 // 队列的最大消息数
	// args["x-max-length-bytes"] = 1048576 // 队列的最字节长度
	args["x-max-priority"] = 10 // 优先级队列
	queue, err := r.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		log.Printf("Queue %s 创建失败: %s\n", name, err)
		return err
	}

	r.Queue[name] = &queue
	return nil
}

func (r *RabbitmqInstance) ExclusiveQueueDeclare(name string) (*amqp.Queue, error) {
	if _, ok := r.Queue[name]; ok {
		log.Printf("Queue %s 已存在 \n", name)
		return nil, errors.New("Queue " + name + "已存在")
	}

	durable := false    // 持久化
	autoDelete := false // 不自动删除
	exclusive := true   // 排他队列，其他用户不可见此队列，只能被当前的连接使用，并且在连接关闭后被删除。
	noWait := false     // 不需要等待应答，默认是需要等待

	args := amqp.Table{}
	// args["x-max-length"] = 1 // 队列的最大消息数
	// args["x-max-length-bytes"] = 1048576 // 队列的最字节长度
	queue, err := r.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		log.Printf("Queue %s 创建失败: %s\n", name, err)
		return nil, err
	}

	r.Queue[queue.Name] = &queue
	return &queue, nil
}

func (r *RabbitmqInstance) QueueDeclareBindDeadExchange(name string, args amqp.Table) error {
	if _, ok := r.Queue[name]; ok {
		log.Printf("Queue %s 已存在 \n", name)
		return errors.New("Queue " + name + "已存在")
	}

	durable := true     // 持久化
	autoDelete := false // 不自动删除
	exclusive := false  // 排他，其他用户不可见此队列，断开连接自动删除
	noWait := false     // 不需要等待应答，默认是需要等待

	queue, err := r.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		log.Printf("Queue %s 创建失败: %s\n", name, err)
		return err
	}

	// 可能是rabbitmq自动生成的随机队列名，所以需要回写
	r.Queue[queue.Name] = &queue
	return nil
}

func (r *RabbitmqInstance) SetPublishHeaders(headers amqp.Table) {
	r.Headers = headers
}

func (r *RabbitmqInstance) SetPublishQueueHeaders(headers amqp.Table) {
	r.QueueHeaders = headers
}

func (r *RabbitmqInstance) Publish(body string, exchangeName, routingKey string) error {
	// mandatory 1.false，找不到合适的queue直接丢弃 2.true 不存在返回给生产者
	// immediate 1.false，没有消费者就返回给生产者决定 2.true 没有消费者就不放到队列
	// mandatory标志告诉服务器至少将该消息route到一个队列中，否则将消息返还给生产者；
	//immediate标志告诉服务器如果该消息关联的queue上有消费者，则马上将消息投递给它，如果所有queue都没有消费者，直接把消息返还给生产者，不用将消息入队列等待消费者了。

	publishing := amqp.Publishing{
		Headers:      r.Headers,
		ContentType:  "text/plain",
		Body:         []byte(body),
		DeliveryMode: amqp.Persistent, // 1=non-persistent, 2=persistent 消息要持久化
		Priority:     0,               // 0-9 优先级队列的时候有用
		Timestamp:    time.Now(),
	}

	if err := r.Channel.Publish(
		exchangeName, // exchange名称
		routingKey,   // routing key
		true,         // mandatory
		false,        // immediate,鉴于RabbitMQ3.0不再支持immediate标志
		publishing,
	); err != nil {
		log.Printf("消息发送失败: %s \n", err)
		return err
	}

	// tracing插件记载的时间戳不对，可能是个BUG
	//fmt.Println("时间:", publishing.Timestamp.Format("2006-01-02 15:04:05"))
	log.Printf("消息发送成功【success】: %s  %s  %s \n", exchangeName, routingKey, body)
	return nil
}

func (r *RabbitmqInstance) PublishPriority(body string, exchangeName, routingKey string, Priority uint8) error {
	// mandatory 1.false，找不到合适的queue直接丢弃 2.true 不存在返回给生产者
	// immediate 1.false，没有消费者就返回给生产者决定 2.true 没有消费者就不放到队列
	// mandatory标志告诉服务器至少将该消息route到一个队列中，否则将消息返还给生产者；
	//immediate标志告诉服务器如果该消息关联的queue上有消费者，则马上将消息投递给它，如果所有queue都没有消费者，直接把消息返还给生产者，不用将消息入队列等待消费者了。

	publishing := amqp.Publishing{
		Headers:      r.Headers,
		ContentType:  "text/plain",
		Body:         []byte(body),
		DeliveryMode: amqp.Persistent, // 1=non-persistent, 2=persistent 消息要持久化
		Priority:     Priority,        // 0-9 优先级队列的时候有用,数字越大，优先级越高
		Timestamp:    time.Now(),
	}

	if err := r.Channel.Publish(
		exchangeName, // exchange名称
		routingKey,   // routing key
		true,         // mandatory
		false,        // immediate,鉴于RabbitMQ3.0不再支持immediate标志
		publishing,
	); err != nil {
		log.Printf("消息发送失败: %s \n", err)
		return err
	}

	// tracing插件记载的时间戳不对，可能是个BUG
	//fmt.Println("时间:", publishing.Timestamp.Format("2006-01-02 15:04:05"))
	log.Printf("消息发送成功【success】: %s  %s  %s \n", exchangeName, routingKey, body)
	return nil
}

func (r *RabbitmqInstance) PublishExpiration(body string, exchangeName, routingKey string, Expiration string) error {
	// mandatory 1.false，找不到合适的queue直接丢弃 2.true 不存在返回给生产者
	// immediate 1.false，没有消费者就返回给生产者决定 2.true 没有消费者就不放到队列
	// mandatory标志告诉服务器至少将该消息route到一个队列中，否则将消息返还给生产者；
	//immediate标志告诉服务器如果该消息关联的queue上有消费者，则马上将消息投递给它，如果所有queue都没有消费者，直接把消息返还给生产者，不用将消息入队列等待消费者了。

	// Expiration 代表消息再队列中被丢弃前能存活的时间
	if err := r.Channel.Publish(
		exchangeName, // exchange名称
		routingKey,   // routing key
		true,         // mandatory
		false,        // immediate,鉴于RabbitMQ3.0不再支持immediate标志
		amqp.Publishing{
			Expiration:   Expiration,
			Headers:      r.Headers,
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent, // 1=non-persistent, 2=persistent 消息要持久化
			Priority:     0,               // 0-9 优先级队列的时候有用
		},
	); err != nil {
		log.Printf("消息发送失败: %s \n", err)
		return err
	}

	log.Printf("消息发送成功【success】: %s  %s  %s \n", exchangeName, routingKey, body)
	return nil
}

func (r *RabbitmqInstance) PublishDelay(body string, exchangeName, routingKey string, headers amqp.Table) error {
	// mandatory 1.false，找不到合适的queue直接丢弃 2.true 不存在返回给生产者
	// immediate 1.false，没有消费者就返回给生产者决定 2.true 没有消费者就不放到队列
	// mandatory标志告诉服务器至少将该消息route到一个队列中，否则将消息返还给生产者；
	//immediate标志告诉服务器如果该消息关联的queue上有消费者，则马上将消息投递给它，如果所有queue都没有消费者，直接把消息返还给生产者，不用将消息入队列等待消费者了。
	if err := r.Channel.Publish(
		exchangeName, // exchange名称
		routingKey,   // routing key
		true,         // mandatory
		false,        // immediate,鉴于RabbitMQ3.0不再支持immediate标志
		amqp.Publishing{
			Headers:      headers,
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent, // 1=non-persistent, 2=persistent 消息要持久化
			Priority:     0,               // 0-9 优先级队列的时候有用
		},
	); err != nil {
		log.Printf("消息发送失败: %s \n", err)
		return err
	}

	log.Printf("消息发送成功【success】: %s  %s  %s \n", exchangeName, routingKey, body)
	return nil
}

func (r *RabbitmqInstance) NotifyPublish(body string, exchangeName, routingKey string) error {

	if err := r.Channel.Publish(
		exchangeName, // exchange名称
		routingKey,   // routing key
		true,         // mandatory
		false,        // immediate,鉴于RabbitMQ3.0不再支持immediate标志
		amqp.Publishing{
			Headers:      r.Headers,
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent, // 1=non-persistent, 2=persistent 消息要持久化
			Priority:     0,               // 0-9 优先级队列的时候有用
		},
	); err != nil {
		log.Printf("消息发送失败: %s \n", err)
		return err
	}

	log.Printf("消息发送成功【success】: %s  %s  %s \n", exchangeName, routingKey, body)
	return nil
}

func (r *RabbitmqInstance) QueueBind(exchangeName, queueName, routingKey string) error {
	// 相同的routing key可以绑定到不同的queue
	nowait := false
	if err := r.Channel.QueueBind(queueName, routingKey, exchangeName, nowait, r.QueueHeaders); err != nil {
		log.Printf("Bind 失败: %s %s %s %s \n", exchangeName, queueName, routingKey, err)
		return err
	}

	return nil
}

func (r *RabbitmqInstance) Close() {
	_ = r.Channel.Close()
	_ = r.Conn.Close()
}
