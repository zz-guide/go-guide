package main

import (
	"log"
	"sync"
	"time"
)

/*
*
1. for range 读取chan,用来替代某些情况下的for select
2. 只能读取到值，不能读取到chan状态
3.使用for-range读取channel，这样既安全又便利，当channel关闭时，for循环会自动退出，
无需主动监测channel是否关闭，可以防止读取已经关闭的channel，造成读到数据为通道所存储的数据类型的零值。
*/
var wg sync.WaitGroup

func main() {
	log.Println("====== START ======")
	var chan1 = make(chan string, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()

		chan1 <- "123"
		chan1 <- "456"
		chan1 <- "789"
		close(chan1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range chan1 {
			log.Printf("v=%v \n", v)
			time.Sleep(1 * time.Second)
		}

		log.Println("...waiting...")
	}()

	wg.Wait()
	log.Println("====== END ======")
}
