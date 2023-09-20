package main

import (
	"log"
	"time"
)

func main() {
	T2()
}

func T1() {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	log.Println("start....")
	ch <- true // 无缓冲，发送方阻塞直到接收方接收到数据。
	log.Println("end....")
	log.Printf("cost %.1f s6\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}

func T2() {
	// 无缓冲其实就是默认缓冲设置为0，因为是int零值
	// 无缓冲
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
		log.Println("ch读取到了")
	}()

	// 缓冲设置为0
	ch1 := make(chan bool, 0)
	go func() {
		time.Sleep(time.Second * 4)
		<-ch1
		log.Println("ch1读取到了")
	}()

	log.Println("start....")
	ch <- true  // 无缓冲，发送方阻塞直到接收方接收到数据。
	ch1 <- true // 无缓冲，发送方阻塞直到接收方接收到数据。
	time.Sleep(time.Second * 10)
}
