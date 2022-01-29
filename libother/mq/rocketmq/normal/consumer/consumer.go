package main

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"go-guide/libother/mq/rocketmq/common"
	"go-guide/libother/mq/rocketmq/config"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	log.Println("启动消费")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	ReceivePushMessage()
	//ReceivePullMessage()
	<-quit
	log.Println("关闭消费")
}

func ReceivePushMessage() {

	conf := config.RocketMqConsumer{
		RocketMqConfig: config.RocketMqConfig{
			Host:       []string{"127.0.0.1:9876"},
			RetryTimes: 3,
			GroupName:  groupName,
		},
		Topic: topic,
	}

	con, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.GroupName),
		consumer.WithNameServer(conf.Host),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: conf.AccessKey,
			SecretKey: conf.SecretKey,
		}),
		consumer.WithPullBatchSize(10),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset), // 选择消费时间(首次/当前/根据时间)
		consumer.WithConsumerModel(consumer.BroadCasting),             // 消费模式(集群消费:消费完其他人不能再读取/广播消费：所有人都能读)
	)

	if err != nil {
		panic(err)
	}

	go func() {
		defer func() {
			log.Println("关闭连接")
			common.EchoError(con.Shutdown())
		}()
		var callback func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error)
		callback = func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			log.Println("开始接收")
			orderlyCtx, _ := primitive.GetOrderlyCtx(ctx)
			log.Printf("orderly context: %v\n", orderlyCtx)
			for i := range msgs {
				time.Sleep(time.Millisecond * 100)
				log.Printf("接收到消息: QueueId:%v, QueueOffset:%v, message:%s, store_host: %v, cur_time: %v\n", msgs[i].Queue.QueueId, msgs[i].QueueOffset, msgs[i].Body, msgs[i].StoreHost, common.NowTimeString())
			}

			return consumer.ConsumeSuccess, nil
		}

		err = con.Subscribe(conf.Topic, consumer.MessageSelector{
			Type:       consumer.TAG,
			Expression: "*",
		}, callback)
		if err != nil {
			common.EchoError(err)
		}

		err = con.Start()
		if err != nil {
			common.EchoError(err)
		}

		select {}
	}()
}

func ReceivePullMessage() {

	// 现在版本不支持
	conf := config.RocketMqConsumer{
		RocketMqConfig: config.RocketMqConfig{
			Host:       []string{"127.0.0.1:9876"},
			RetryTimes: 3,
			GroupName:  groupName,
		},
		Topic: topic,
	}

	con, err := rocketmq.NewPullConsumer(
		consumer.WithGroupName(conf.GroupName),
		consumer.WithNameServer(conf.Host),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: conf.AccessKey,
			SecretKey: conf.SecretKey,
		}),
		//consumer.WithPullBatchSize(10),
		//consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset), // 选择消费时间(首次/当前/根据时间)
		//consumer.WithConsumerModel(consumer.BroadCasting),             // 消费模式(集群消费:消费完其他人不能再读取/广播消费：所有人都能读)
	)

	if err != nil {
		panic(err)
	}

	err = con.Start()
	if err != nil {
		common.EchoError(err)
	}

	queue := primitive.MessageQueue{
		Topic:      topic,
		BrokerName: "broker1", // 使用broker的名称
		QueueId:    0,
	}

	err = con.Shutdown()
	if err != nil {
		log.Println("Shutdown Pull Consumer error: ", err)
	}

	offset := int64(0)
	ctx := context.Background()
	for {
		resp, err := con.PullFrom(ctx, queue, offset, 10)
		if err != nil {
			if err == rocketmq.ErrRequestTimeout {
				log.Printf("timeout\n")
				time.Sleep(time.Second)
				continue
			}
			log.Printf("unexpected error: %v\n", err)
			return
		}

		if resp.Status == primitive.PullFound {
			log.Printf("成功拉取到数据. nextOffset: %d\n", resp.NextBeginOffset)
			for _, ext := range resp.GetMessageExts() {
				log.Printf("pull msg: %s\n", ext)
			}
		}

		offset = resp.NextBeginOffset
	}
}
