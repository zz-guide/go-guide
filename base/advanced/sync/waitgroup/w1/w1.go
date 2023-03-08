package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

/**
sync.WaitGroup 用来等待一组goroutine全部执行完成
1.sync.WaitGroup源码分析
	https://zhuanlan.zhihu.com/p/365288361

2.state1字段一共12个字节，分为3部分，计数器数量+goroutine数量+sema信号量


3.使用过程需要注意
	1.ADD需要再wait之前执行，不要包裹在go func中
	2.只要计数器为0就停止执行了，所以要确保先添加所有的任务
	3.不能复制
	4.内部panic无法感知
*/

func main() {
	//T1()
	T2()
}

func T1() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(p int) {
			if p%2 == 0 {
				time.Sleep(time.Second * 5)
			}
			wg.Add(1)
			log.Println("p:", p)
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println("asdasdasd")
}

func T2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		log.Println("协程1")
		wg.Done()
	}()

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover:", err)
			}
		}()
		panic("协程2 panic")
		log.Println("协程2")
	}()

	go func() {
		log.Println("协程3")
		wg.Done()
	}()
	wg.Wait()
	log.Println("asdasdasd")
}
