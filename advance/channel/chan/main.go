package main

import (
	"log"
	"sync"
)

/*
*
结论:
1. 通过make函数创建chan
2. chan 必须是多个不同协程之间进行同行，接收端和发送端不能处于同一个协程内部
3. sync.WaitGroup 用来等待协程全部执行完毕
4. chan必须初始化才可使用，否则会死锁：all goroutines are asleep - deadlock!chan send (nil chan)
*/

var wg sync.WaitGroup

func main() {
	log.Println("====== START ======")
	var chan1 = make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		chan1 <- "hello world"
		close(chan1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		v, ok := <-chan1
		log.Printf("v=%v, ok=%v\n", v, ok)
	}()

	wg.Wait()

	log.Println("====== END ======")
}
