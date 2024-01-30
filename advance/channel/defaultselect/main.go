package main

import (
	"fmt"
	"time"
)

/*
*
1. 只有default的select，跟串行化没区别
*/

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("没有内容的select 会阻塞")
	}()
	select {
	default:
		fmt.Println("======")
	}
}
