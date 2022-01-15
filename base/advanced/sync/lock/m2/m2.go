package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Mutex原理：
type Mutex struct {
	state int32
	sema  uint32
}

//正常模式：正常模式下所有的goroutine按照FIFO的顺序进行锁获取,被唤醒的goroutine和新请求锁的goroutine同时进行锁获取，通常新请求锁的goroutine更容易获取锁；不公平
//饥饿模式：饥饿模式所有尝试获取锁的goroutine进行等待排队，新请求锁的goroutine不会进行锁获取，而是加入队列尾部等待获取锁；公平

Lock的过程：1.先原子获取锁一次；2.不成功继续自旋；3。还不成功就加入队列等待唤醒。

自旋意义：
1、GOMAXPROCS >1
2. 多核
3.至少有一个P正在running
4.当前P的本地队列是空的

1.runtime_doSpin 底层：30 PAUSE procyield 4次
2.结束自旋的条件：超过4次，锁释放，饥饿模式
3.


锁被释放以后排队的g需要和不排队的竞争，自旋的更有优势，本身就在cpu上运行，自旋的有很多，被唤醒的每次只有一个，被唤醒的g插到头部。
加锁等待时间超过1ms，模式变为饥饿模式

*/

// 下边的变量共用一个表达式
const (
	// state 字段
	mutexLocked   = 1 << iota // mutex is locked  第一位
	mutexWoken                // 是否有g已唤醒，1是；0-否 第二位
	mutexStarving             // 模式：0-正常模式 1-饥饿；4  第三位
	aa
	bb
	mutexWaiterShift = iota // 记录等待排队的g

	// 加锁等待时间 1ms
	starvationThresholdNs = 1e6
)

func main() {
	//F1()
	fmt.Println("mutexLocked:", mutexLocked)
	fmt.Println("mutexWoken:", mutexWoken)
	fmt.Println("mutexStarving:", mutexStarving)
	fmt.Println("aa:", aa)
	fmt.Println("bb:", bb)
	fmt.Println("mutexWaiterShift:", mutexWaiterShift)
}

// F1
// 结论：
//*
func F1() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var count int64

	t := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			count++
		}(&wg)
	}

	wg.Wait()
	fmt.Printf("test1 花费时间:%d, count的值为:%d \n", time.Now().Sub(t), count)
}
