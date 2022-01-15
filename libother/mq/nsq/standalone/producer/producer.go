package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

var topic = "bingo"

func main() {
	//SendMessage()
	RandomSendMessage()
}

func CreateProducer(addr string) (*nsq.Producer, error) {
	var producer *nsq.Producer
	var err error
	producer, err = nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		fmt.Println("创建生产者失败...:", err)
		return producer, err
	}

	err = producer.Ping()
	if nil != err {
		// 关闭生产者
		producer.Stop()
		producer = nil
		fmt.Println("Ping失败...:", err)
		return producer, err
	}

	return producer, err
}

func RandomSendMessage() {
	producer1, err := CreateProducer("localhost:4150")
	producer2, err := CreateProducer("localhost:5150")
	if err != nil {
		fmt.Println("创建生产者失败...:", err)
		return
	}

	if producer1 == nil {
		fmt.Println("producer1 nil...")
		return
	}

	if producer2 == nil {
		fmt.Println("producer2 nil...")
		return
	}

	for i := 0; i < 1; i++ {
		message := fmt.Sprintf("{'name':'许磊','id':'%d','msg':'123'}", i)

		if i%2 == 0 {
			err = producer1.Publish(topic, []byte(message))
			if err != nil {
				fmt.Printf("producer.Publish,err : %v", err)
			}

			fmt.Printf("msg %d 成功(生产者1): %s \n", i, message)
		} else {
			err = producer2.Publish(topic, []byte(message))
			if err != nil {
				fmt.Printf("producer.Publish,err : %v", err)
			}

			fmt.Printf("msg %d 成功(生产者2): %s \n", i, message)
		}

	}

	fmt.Println("全部发送完毕")
}

func SendMessage() {
	producer1, err := CreateProducer("localhost:4150")
	if err != nil {
		fmt.Println("创建生产者失败...:", err)
		return
	}

	if producer1 == nil {
		fmt.Println("producer1 nil...")
		return
	}

	for i := 0; i < 10000; i++ {
		message := fmt.Sprintf("{'name':'许磊','id':'%d','msg':'123'}", i)
		err = producer1.Publish(topic, []byte(message))
		if err != nil {
			fmt.Printf("producer.Publish,err : %v", err)
		}

		fmt.Printf("message %d 发送成功: %s \n", i, message)
	}

	fmt.Println("全部发送完毕")
}
