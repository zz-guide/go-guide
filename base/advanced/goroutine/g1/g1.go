package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
	结论：
		1.通过runtime.NumGoroutine()可以打印出g的数量
 */
func main() {
	go func() {
		go func() {
			fmt.Println("g的数量:", runtime.NumGoroutine())
		}()

		time.Sleep(time.Second * 5)
	}()

	time.Sleep(time.Second * 5)
}
