package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

var rwMutex sync.RWMutex

func runReadLock(q int) {
	log.Printf("%d 准备获取读锁\n", q)
	rwMutex.RLock() // 与写锁构成互斥，在读的时候不允许写
	defer func() {
		rwMutex.RUnlock()
		log.Printf("%d 读锁解除\n", q)
	}()

	log.Printf("%d 持有读锁，运行中。。。\n", q)
	time.Sleep(time.Second * 2)
}

func runWriteLock(q int) {
	log.Printf("%d 准备获取写锁\n", q)
	rwMutex.Lock()
	defer func() {
		rwMutex.Unlock()
		log.Printf("%d 写锁解除\n", q)
	}()

	log.Printf("%d 持有写锁，运行中。。。\n", q)
	time.Sleep(time.Second * 4)
}

func main() {
	runtime.GOMAXPROCS(10)
	for q := 0; q < 2; q++ {
		go func(q int) {
			runWriteLock(q)
		}(q)
	}

	for k := 0; k < 3; k++ {
		go func(k int) {
			runReadLock(k)
		}(k)
	}

	time.Sleep(time.Second * 30)
}
