package main

import (
	"fmt"
	"log"
	"time"
)

/*
*
1. 空的select 会阻塞
*/

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("没有内容的select 会阻塞")
	}()
	log.Println("========")
	// 一直阻塞 fatal error: all goroutines are asleep - deadlock!
	select {}
}
