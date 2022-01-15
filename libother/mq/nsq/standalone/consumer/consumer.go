package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
)

var topic = "bingo"
var address = "192.168.3.200:4161"

func main() {
	InitConsumer("ch1")

	fmt.Println("消费者启动成功")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("main主动退出")
}

func InitConsumer(channel string) {
	c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		fmt.Println("消费者创建失败")
		panic(err)
	}

	if c == nil {
		fmt.Println("消费者 nil")
		return
	}

	c.SetLogger(nil, 0)        //屏蔽系统日志
	c.AddHandler(&ConsumerT{}) // 添加消费者接口

	// 连接到nsqlookup或者nsq
	if err := c.ConnectToNSQLookupd(address); err != nil {
		fmt.Println("消费者连接失败")
		panic(err)
	}

	fmt.Println("sss:", c.Stats())
}

type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	//return errors.New("消费异常")
	os.Exit(-1) // 模拟异常退出
	return nil
}
