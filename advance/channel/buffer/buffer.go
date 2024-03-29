package main

import (
	"log"
	"sync"
	"time"
)

/*
*
1. 无缓冲，发送方阻塞直到接收方接收到数据。
2. make(chan string, 0) 设置为0，或者不写，都是无缓冲的，只能发送一个接收一个
*/
var wg sync.WaitGroup

func main() {
	log.Println("====== START ======")
	var chan1 = make(chan string, 1)
	// 发送方
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("===立即发送===")
		chan1 <- "123"
		log.Println("===发送完毕 1===") // 此处会立即开始打印，缓冲未满

		chan1 <- "234"
		log.Println("===发送完毕 2===") // 此处会阻塞，缓冲已满
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

		v1, ok1 := <-chan1
		log.Printf("接收数据：v1=%v, ok1=%v\n", v1, ok1)

	}()

	wg.Wait()
	log.Println("====== END ======")
}
