package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	fmt.Println("start....")
	ch <- true // 无缓冲，发送方阻塞直到接收方接收到数据。
	fmt.Println("end....")
	fmt.Printf("cost %.1f s6\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}
