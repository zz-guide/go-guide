package main

import (
	"sync"
	"time"
)

/**
RWMutex（读写锁）
1、RWMutex 是单写多读锁，该锁可以加多个读锁或者一个写锁
2、读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁
3、写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占适用于读多写少的场景
4. 声明锁变量的时候不要使用指针，当成函数参数传递的时候需要使用指针
*/
func main() {
	go read(1)
	go read(2)

	time.Sleep(2 * time.Second)
}

var m sync.RWMutex

func read(i int) {
	println(i, "read start")

	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over")
}
