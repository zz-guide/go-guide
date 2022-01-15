package main

import (
	"github.com/beanstalkd/go-beanstalk"
	"log"
	"time"
)

var (
	tube  = "tube1"
	tube2 = "tube2"

	ttr = time.Second * 15 // 10秒内没响应状态，重新入队列，可能会被同一个消费者取多次
)

func main() {
	// 只有当消费者处理能力不足的时候优先级队列才有意义，否则就是谁先到谁先处理
	for i := 0; i < 10; i++ {
		SendMessage()
		SendMessage1()
	}
}

func SendMessage() {
	conn, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		log.Println("connect 出错:", err)
		return
	}

	// 创建一个tube
	tube := beanstalk.NewTube(conn, tube)

	body := "{'name':'低','tube':'1'}"

	// 往tube发送job，

	id, err := tube.Put([]byte(body), 42, 5*time.Second, ttr)
	if err != nil {
		log.Printf("put失败: %s \n", err)
		return
	}

	newBody, err := conn.Peek(id) // peek用来查看job，不改变job状态
	if err != nil {
		log.Printf("peek失败: %s \n", err)
		return
	}

	log.Println("发送成功:", id, string(newBody))
}

func SendMessage1() {
	conn, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		log.Println("connect 出错:", err)
		return
	}

	// 创建一个tube
	tube := beanstalk.NewTube(conn, tube)

	body := "{'name':'高','tube':'1'}"

	// 往tube发送job，

	id, err := tube.Put([]byte(body), 21, 5*time.Second, ttr)
	if err != nil {
		log.Printf("put失败: %s \n", err)
		return
	}

	newBody, err := conn.Peek(id) // peek用来查看job，不改变job状态
	if err != nil {
		log.Printf("peek失败: %s \n", err)
		return
	}

	log.Println("发送成功:", id, string(newBody))
}
