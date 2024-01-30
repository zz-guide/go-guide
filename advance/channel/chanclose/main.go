package main

import (
	"log"
	"sync"
)

/*
*
结论:
1. close之后不能再发送数据, 会报错 panic: send on closed channel
2. v,ok := <-chan1 ,channel读取有2个变量，一个代表通道里的值，一个代表chan是否被close
3. 即使被close，仍然可以从chan读取数据，只不过读取到的是对应类型的“零值”
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
		// chan1 <- "closed..."
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		v, ok := <-chan1
		log.Printf("v=%v, ok=%v\n", v, ok)

		// 再次读取，返回 “”,false
		v1, ok1 := <-chan1
		log.Printf("v1=%v, ok1=%v\n", v1, ok1)
	}()

	wg.Wait()

	log.Println("====== END ======")
}
