package main

import (
	"log"
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
			log.Println("g的数量:", runtime.NumGoroutine())
		}()

		time.Sleep(time.Second * 5)
	}()

	time.Sleep(time.Second * 5)
}
