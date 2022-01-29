package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"go-guide/libother/mq/rocketmq/common"
	"go-guide/libother/mq/rocketmq/config"
	"log"
	"os"
	"time"

	. "go-guide/libother/mq/rocketmq/common"
)

var (
	topic     = "bingo"
	groupName = "bingo-group"
)

func init() {
	rlog.SetLogLevel("error")
}

func main() {
	SendSyncMessage()
	//SendAsyncMessage()
}

func SendSyncMessage() {
	conf := config.RocketMqProducer{
		RocketMqConfig: config.RocketMqConfig{
			Host:       []string{"127.0.0.1:9876"},
			RetryTimes: 3,
			GroupName:  groupName,
		},
		SendMsgTimeout: 5 * time.Second,
		Topic:          topic,
	}

	p, err := rocketmq.NewProducer(
		producer.WithSendMsgTimeout(conf.SendMsgTimeout),
		producer.WithGroupName(conf.GroupName),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: conf.AccessKey,
			SecretKey: conf.SecretKey,
		}),
		producer.WithNameServer(conf.Host),
		producer.WithRetry(conf.RetryTimes),
		producer.WithCreateTopicKey(conf.Topic),
	)
	if err != nil {
		panic(err)
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	defer func() {
		common.EchoError(p.Shutdown())
	}()

	// 批量发送消息
	var messages []*primitive.Message
	messages = append(messages, primitive.NewMessage(conf.Topic, GetMessage("sync message")))

	// 发送带有TAG的消息
	msg1 := primitive.NewMessage(conf.Topic, GetMessage("A Tag"))
	msg1.WithTag("tagA")
	messages = append(messages, msg1)
	// 发送同步消息
	res, err := p.SendSync(context.Background(), messages...)
	if err != nil {
		log.Printf("发送失败 : %s\n", err)
		return
	}

	fmt.Printf("发送成功: result=%s\n", res.String())
}

func SendAsyncMessage() {
	conf := config.RocketMqProducer{
		RocketMqConfig: config.RocketMqConfig{
			Host:       []string{"127.0.0.1:9876"},
			RetryTimes: 3,
			GroupName:  groupName,
		},
		SendMsgTimeout: 5 * time.Second,
		Topic:          topic,
	}

	p, err := rocketmq.NewProducer(
		producer.WithSendMsgTimeout(conf.SendMsgTimeout),
		producer.WithGroupName(conf.GroupName),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: conf.AccessKey,
			SecretKey: conf.SecretKey,
		}),
		producer.WithNameServer(conf.Host),
		producer.WithRetry(conf.RetryTimes),
		producer.WithCreateTopicKey(conf.Topic),
	)
	if err != nil {
		panic(err)
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	defer func() {
		common.EchoError(p.Shutdown())
	}()

	msg := &primitive.Message{
		Topic: conf.Topic,
		Body:  GetMessage("异步 message"),
	}

	// 发送异步消息
	err = p.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			fmt.Printf("异步接收消息错误:%v\n", err)
		} else {
			fmt.Printf("异步发送消息成功. result=%s\n", result.String())
		}
	}, msg)
	if err != nil {
		log.Printf("发送失败 : %s\n", err)
		return
	}
	time.Sleep(time.Second * 2)
}
