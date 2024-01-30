package main

import (
	"log"
	"sync"
	"time"
)

/*
*
1. for select 可组合读取通道数据
2. select 如果没定义default， select 将阻塞，直到某个通信可以运行
3. select 如果定义了default, 执行default
4. 使用goto 或者 break + 标记的方式可以跳出循环
5. default 可以当做是chan没准备好接收之前，为了不阻塞而设置，一旦准备好不会再进入到default(可能以break)
6. select 匹配无序
7. select默认只匹配一次
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
	AAA:
		for {
			select {
			case v, ok := <-chan1:
				if !ok {
					break AAA
				}

				log.Printf("v=%v, ok=%v\n", v, ok)
				time.Sleep(1 * time.Second)
			default:
				log.Println("default")
			}

			log.Println("...waiting...")
		}
	}()

	wg.Wait()
	log.Println("====== END ======")
}
