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
	"os"
	"time"
)

var (
	topic     = "bingo"
	groupName = "bingo-group"
)

func init() {
	rlog.SetLogLevel("error")
}

func main() {
	SendDelaySyncMessage()
}

func SendDelaySyncMessage() {
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

	// 发送延迟消息
	msg := primitive.NewMessage(topic, []byte("延迟同步消息"))
	msg.WithDelayTimeLevel(2)
	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("延迟同步消息错误:%s\n", err)
	} else {
		fmt.Printf("延迟同步消息成功. result=%s\n", res.String())
	}

	fmt.Printf("发送成功: result=%s\n", res.String())
}
