package main

import (
	"fmt"
	"sync"
	"time"
)

/**
tips：什么是原子操作？
简单的说cpu能在一个指令中完成的操作就可以称为原子操作，例如64位的操作系统，一次可以处理8byte的内容，如果一次赋值的内容小于等于8byte，那么可以认为这是一个single machine word，是由系统底层保证操作的原子性。
既然我们看到了汇编指令，就引出了cpu的指令重排，cpu的设计者为了更加充分的利用CPU的性能，例如分支预测，流水线等，指令重排会将读写指令进行重排，但是会保证最后结果的正确，但是在多线程的环境下就会出问题，造成数据不正确。来看一看内存模型

我们知道CPU为了平衡内核，内存，硬板之间的速度差异，有了3级缓存的策略，从一级缓存到3级缓存速度由快到慢，离cpu越远，速度就越慢，主要目的还是为了减少CPU访问主内存的次数。那么，试想图中的两个线程同时 操作 A，B两个共享变量，他们先将各自操作的变量存储在store buffer，然后互相再去访问对方的变量，这个时候各自的变量还没有刷新到内存，结果拿到的都是互相的初始值，这个时候再去那这个初始值去操作就会出问题，所以在多线程的环境下CPU提供了 barrier指令，要求所有对内存的操作需要扩散到主内存之后，才能进行其他的内存的操作。

go run -race m2.go
==================
test1 花费时间：6813987, count的值为：100
Found 1 data race(s6)
exit status 66

或者在使用锁的过程中复制了锁，例如函数的代码调用，当做参数传过去，重新进行加锁，解锁就会造成意想不到的结果，
因为锁是有状态的，复制锁的时候会将锁的状态一起复制过去。对于这种复制锁造成的问题，
可以使用go vet 来检查代码中的锁复制问题

golang不提供重入锁，不是必须的，可通过更改代码结构来避免

Mutex（互斥锁）
1、Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
2、在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
3、使用 Lock() 加锁后，不能再继续对其加锁，直到 Unlock() 解锁后才能再加锁
4、在 Lock() 之前使用 Unlock() 会导致 panic 异常
5、已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
6、在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁适用于读写不确定，并且只有一个读或者写的场景
*/
func main() {
	F1()
}

// F1
// 结论：
//	1.同一个goroutine不能多次Lock
//	2.fatal error: all goroutines are asleep - deadlock!
//	2.fatal error: fatal error out of memory
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
			//defer mutex.Unlock()
			mutex.Lock()
			// 不是原子操作
			// 1.先把count从内存读取到寄存器
			// 2.寄存器执行加1
			// 3.再把值写到内存
			count++
		}(&wg)
	}

	wg.Wait()
	fmt.Printf("test1 花费时间：%d, count的值为：%d \n", time.Now().Sub(t), count)
}
