package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
结论:
	CPU如何实现原子操作
	CPU 处理器速度远远大于在主内存中的，为了解决速度差异，在他们之间架设了多级缓存，如 L1、L2、L3 级别的缓存，这些缓存离CPU越近就越快，将频繁操作的数据缓存到这里，加快访问速度 ，如下图所示：
	现在都是多核 CPU 处理器，每个 CPU 处理器内维护了一块字节的内存，每个内核内部维护着一块字节的缓存，当多线程并发读写时，就会出现缓存数据不一致的情况。
	此时，处理器提供：
	1.总线锁定
		当一个处理器要操作共享变量时，在 BUS 总线上发出一个 Lock 信号，其他处理就无法操作这个共享变量了。
		缺点很明显，总线锁定在阻塞其它处理器获取该共享变量的操作请求时，也可能会导致大量阻塞，从而增加系统的性能开销。
	2.缓存锁定
		后来的处理器都提供了缓存锁定机制，也就说当某个处理器对缓存中的共享变量进行了操作，其他处理器会有个嗅探机制，将其他处理器的该共享变量的缓存失效，待其他线程读取时会重新从主内存中读取最新的数据，基于 MESI 缓存一致性协议来实现的。
		现代的处理器基本都支持和使用的缓存锁定机制。

	注意：
	有如下两种情况处理器不会使用缓存锁定：
	（1）当操作的数据跨多个缓存行，或没被缓存在处理器内部，则处理器会使用总线锁定。
（2）有些处理器不支持缓存锁定，比如：Intel 486 和 Pentium 处理器也会调用总线锁定

	CAS底层指令
	底层硬件通过将 CAS 里的多个操作在硬件层面语义实现上，通过一条处理器指令保证了原子性操作。这些指令如下所示：
	（1）测试并设置（Tetst-and-Set）
	（2）获取并增加（Fetch-and-Increment）
	（3）交换（Swap）
	（4）比较并交换（Compare-and-Swap）
	（5）加载链接/条件存储（Load-Linked/Store-Conditional）
	前面三条大部分处理器已经实现，后面的两条是现代处理器当中新增加的。而且根据不同的体系结构，指令存在着明显差异。
	在IA64，x86 指令集中有 cmpxchg 指令完成 CAS 功能，在 sparc-TSO 也有 casa 指令实现，而在 ARM 和 PowerPC 架构下，则需要使用一对 ldrex/strex 指令来完成 LL/SC 的功能。
	在精简指令集的体系架构中，则通常是靠一对儿指令，如：load and reserve 和 store conditional 实现的，在大多数处理器上 CAS 都是个非常轻量级的操作，这也是其优势所在。

	缺点：
	1.CAS在共享资源竞争比较激烈的时候，每个goroutine会容易处于自旋状态，影响效率，在竞争激烈的时候推荐使用锁。
		加锁失败，自旋会长时间占用 CPU 资源，加大了系统性能开销。
	2.ABA问题是无锁结构实现中常见的一种问题，可基本表述为：
		进程P1读取了一个数值A
		P1被挂起(时间片耗尽、中断等)，进程P2开始执行
		P2修改数值A为数值B，然后又修改回A
		P1被唤醒，比较后发现数值A没有变化，程序继续执行。
		A-->B--->A 问题，假设有一个变量 A ，修改为B，然后又修改为了 A，实际已经修改过了，但 CAS 可能无法感知，造成了不合理的值修改操作。
		整数类型还好，如果是对象引用类型，包含了多个变量，那怎么办？加个版本号或时间戳呗，没问题！

	乐观锁会以一种更加乐观的态度对待事情，认为自己可以操作成功。当多个线程操作同一个共享资源时，仅能有一个线程同一时间获得锁成功，在乐观锁中，其他线程发现自己无法成功获得锁，并不会像悲观锁那样阻塞线程，而是直接返回，可以去选择再次重试获得锁，也可以直接退出。
*/
func main() {
	CASExample()
}

// AddExample
//AddInt32可以实现对元素的原子增加或减少，函数会直接在传递的地址上进行修改操作。
//addr需要修改的变量的地址，delta修改的差值[正或负数]，返回new修改之后的新值。
//AddInt32(),AddUint32(),AddInt64(),AddUint64(),AddUintptr()
//**
func AddExample() {
	var c int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 多个协程同时对该变量进行+1操作，最终结构可能不是100
			// c++
			// 通过AddInt32保证最终结果就是100
			atomic.AddInt32(&c, 1)
		}(&wg)
	}
	wg.Wait()
	fmt.Println("c:", c)
}

// CASExample
//函数会先判断参数addr指向的值与参数old是否相等，如果相等，则用参数new替换参数addr的值。最后返回swapped是否替换成功。
//CompareAndSwapInt32(),CompareAndSwapInt64(),CompareAndSwapUint32(),CompareAndSwapUint64(),CompareAndSwapUintptr(),CompareAndSwapPointer()
//**
func CASExample() {
	var c int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 需要先原子性读取，然后再通过cas乐观锁赋值
			tmp := atomic.LoadInt32(&c)
			// 替换成功返回true，无法保证c最终是100
			for {
				if atomic.CompareAndSwapInt32(&c, tmp, tmp+1) {
					break
				}
			}
		}(&wg)
	}
	wg.Wait()
	fmt.Println("c : ", c)

	//需要注意的是如果是uint32,unint64时,不能直接传负数，所以需要利用二进制补码机制
	//var b uint32
	//b += 20
	//atomic.AddUint32(&b, ^uint32(10-1)) // 等价于 b -= 10
	// atomic.Adduint32(&b, ^uint32(N-1)) //N为需要减少的正整数值
	//fmt.Println(b == 10) // true
}

// SwapExample
//Swap会直接执行赋值操作，并将原值作为返回值返回
//SwapInt32(),SwapInt64(),SwapUint32(),SwapUint64(),SwapUintptr(),SwapPointer()
//**
func SwapExample() {
	var c int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			tmp := atomic.LoadInt32(&c)
			// tmp+1不是一个原子操作，直接赋值才是原子操作
			// 不能保证最终是100
			atomic.SwapInt32(&c, tmp+1)
		}(&wg)
	}
	wg.Wait()
	fmt.Println("c : ", c)
}

// StoreExample
//StoreExample会直接执行赋值操作
//StoreInt32(),StoreInt64(),StoreUInt32(),StoreUInt64(),StoreUintptr(),StorePointer()
//LoadInt32(),LoadInt64(),LoadUInt64(),LoadUInt64(),LoadUintptr(),LoadPointer()
//**
func StoreExample() {
	var c int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			tmp := atomic.LoadInt32(&c)
			// tmp+1不是一个原子操作，直接赋值才是原子操作
			// 不能保证最终是100
			atomic.StoreInt32(&c, tmp+1)
		}(&wg)
	}
	wg.Wait()
	fmt.Println("c : ", c)
}
