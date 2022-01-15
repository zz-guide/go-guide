package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	ch := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
		fmt.Printf("读取完毕")
	}()
	ch <- true
	ch <- true                                                 // 缓冲区为 2，发送方不阻塞，继续往下执行
	fmt.Printf("cost %.1f s6\n", time.Now().Sub(st).Seconds()) // cost 0.0 s6
	ch <- true                                                 // 缓冲区使用完，发送方阻塞，2s 后接收方接收到数据，释放一个插槽，继续往下执行
	fmt.Printf("cost %.1f s6\n", time.Now().Sub(st).Seconds()) // cost 2.0 s6
	time.Sleep(time.Second * 5)
}
