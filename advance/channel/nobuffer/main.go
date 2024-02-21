package main

import (
	"log"
	"sync"
	"time"
)

/*
*
1. 无缓冲，可以认为无法存储数据，相等于同步读写。
2. make(chan string, 0) 设置为0，或者不写，都是无缓冲的，只能发送一个接收一个
*/
var wg sync.WaitGroup

func main() {
	log.Println("====== START ======")
	var chan1 = make(chan string, 0)
	// 发送方
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("===立即发送===")
		chan1 <- "123"
		log.Println("===发送完毕===") // 无缓冲的话，此处会阻塞，直到开始接收
		close(chan1)
	}()

	// 接收方
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("===延迟接收端===")
		time.Sleep(time.Second * 2)
		v, ok := <-chan1
		log.Printf("接收数据：v=%v, ok=%v\n", v, ok)

	}()

	wg.Wait()
	log.Println("====== END ======")
}
