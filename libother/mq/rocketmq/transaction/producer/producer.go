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
	"sync"
	"sync/atomic"
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
}

type DemoListener struct {
	localTrans       *sync.Map
	transactionIndex int32
}

func NewDemoListener() *DemoListener {
	return &DemoListener{
		localTrans: new(sync.Map),
	}
}

func (dl *DemoListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	nextIndex := atomic.AddInt32(&dl.transactionIndex, 1)
	fmt.Printf("nextIndex: %v for transactionID: %v\n", nextIndex, msg.TransactionId)
	status := nextIndex % 3
	dl.localTrans.Store(msg.TransactionId, primitive.LocalTransactionState(status+1))

	fmt.Printf("dl")
	//在SendMessageInTransaction 方法调用ExecuteLocalTransaction方法，
	//如果ExecuteLocalTransaction 返回primitive.UnknowState 那么brocker就会调用CheckLocalTransaction方法检查消息状态
	// 如果返回  primitive.CommitMessageState 和primitive.RollbackMessageState 则不会调用CheckLocalTransaction
	return primitive.UnknowState
}

func (dl *DemoListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Printf("%v msg transactionID : %v\n", time.Now(), msg.TransactionId)
	v, existed := dl.localTrans.Load(msg.TransactionId)
	if !existed {
		fmt.Printf("unknow msg: %v, return Commit", msg)
		return primitive.CommitMessageState
	}
	state := v.(primitive.LocalTransactionState)
	switch state {
	case 1:
		fmt.Printf("checkLocalTransaction COMMIT_MESSAGE: %v\n", msg)
		return primitive.CommitMessageState
	case 2:
		fmt.Printf("checkLocalTransaction ROLLBACK_MESSAGE: %v\n", msg)
		return primitive.RollbackMessageState
	case 3:
		fmt.Printf("checkLocalTransaction unknow: %v\n", msg)
		return primitive.UnknowState
	default:
		fmt.Printf("checkLocalTransaction default COMMIT_MESSAGE: %v\n", msg)
		return primitive.CommitMessageState
	}
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

	p, err := rocketmq.NewTransactionProducer(
		NewDemoListener(),
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

	ctx := context.Background()
	for _, message := range messages {
		// 发送同步消息
		res, err := p.SendMessageInTransaction(ctx, message)
		if err != nil {
			log.Printf("发送失败 : %s\n", err)
			return
		}

		fmt.Printf("发送成功: result=%s\n", res.String())
	}

}
