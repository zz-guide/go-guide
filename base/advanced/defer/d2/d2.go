package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//deferExit()
	deferPanic()
}

// 结论：1.os.Exit 可导致defer不执行
// defer执行的时机：
//1.包裹defer的函数return
//2.包裹defer的函数执行到末尾
//3.goroutine panic的时候会执行本goroutine的所有defer
func deferExit() {
	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}

func deferPanic() {
	defer func() {
		log.Println("defer")
	}()

	defer func() {
		if err := recover(); err != nil {
			log.Println("外部 err:", err)
		}
	}()

	go func() {
		// 外部goroutine无法recover其他协程的panic
		/*defer func() {
			if err := recover(); err != nil {
				log.Println("err:", err)
			}
		}()*/

		// panic的时候只会执行本goroutine的defer
		panic("出错了")
	}()

	time.Sleep(time.Second * 2)
}
