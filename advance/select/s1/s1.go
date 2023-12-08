package main

import (
	"fmt"
	"time"
)

/**
select的原理：select是Golang在语言层面提供的多路IO复用的机制。与switch语句稍微有点相似，也会有case和最后的default选择支。每一个case代表一个通信操作（在某个channel上进行发送或者接收）并且会包含一些语句组成一个语句块。

select {
	case 表达式1:
		<code>
	case 表达式2:
		<code>
  default:
	  <code>
}

在运行 select 时，会遍历所有（如果有机会的话）的 case 表达式，只要有一个信道有接收到数据，那么 select 就结束

总结
	1.select语句中除default外，各case执行顺序是随机的
	2.select语句中如果没有default语句，则会阻塞等待任意一个case
	3.select语句中读操作要判断是否成功读取，关闭的channel也可以读取
	4.select语句中除default外，每个case只能操作一个channel，要么读要么写
	5.select 与 switch 原理很相似，但它的使用场景更特殊，学习了本篇文章，你需要知道如下几点区别：
	6.select 只能用于 channel 的操作(写入/读出)，而 switch 则更通用一些；
	7.select 的 case 是随机的，而 switch 里的 case 是顺序执行；
	8.select 要注意避免出现死锁，同时也可以自行实现超时机制；
	9.select 里没有类似 switch 里的 fallthrough 的用法；
	10.select 不能像 switch 一样接函数或其他表达式。

注意：
	1.避免造成死锁，select 在执行过程中，必须命中其中的某一分支。
	2.select具有随机性
	3.select超时设置
		当 case 里的信道始终没有接收到数据时，而且也没有 default 语句时，select 整体就会阻塞，但是有时我们并不希望 select 一直阻塞下去，这时候就可以手动设置一个超时时间。使用context
	4.读取/写入都可以，select 里的 case 表达式只要求你是对信道的操作即可，不管你是往信道写入数据，还是从信道读出数据



源码总结:
	1.初始化阶段：确定case的轮询顺序 pollOrder 和加锁顺序 lockOrder
	通过 runtime.fastrandn 函数，打乱case的访问顺序
	下面步骤2访问的时候就是按照乱序访问的，这也就是为什么多个case ready随机执行一个的原因
	加锁顺序lockOrder具体是用一个简单的堆排序对channel地址排序实现，为了在读数据时能顺序给channel加锁，和去重防止对同一个channel多次加锁
	2.主流程
	首先在for循环中对case进行遍历，查看是否ready，已经ready就直接跳到处理部分，流程就结束
	如果有ready的case就随机执行一个，没有的基础上如果有default，执行default语句
	挂起
	通道都没有ready，且没有default语句，所以把当前goroutine挂到每一个通道的等待队列中等待唤醒
	唤醒 有channel准备好了，当前 Goroutine 就会被调度器唤醒，返回当前case，其他case中通道队列移除该goroutine，不再等待
	select 关键字是 Go 语言特有的控制结构，它的实现原理比较复杂，需要编译器和运行时函数的通力合作
*/
func main() {
	blockExample()
}

const (
	fmat = "2006-01-02 15:04:05"
)

func randExample() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c1 <- "hello"
	c2 <- "world"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		fmt.Println("No data received.")
	}
}

// timeoutExample
// 结论: 1.for和select中间是会阻塞的，直到有case可以读写
//		2.for循环select时，如果其中一个case通道已经关闭，则每次都会执行到这个case
//
//**
func timeoutExample() {
	c := make(chan string, 1)
	timeout := time.After(5 * time.Second)

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("开始传输")
		c <- "你好"
		close(c)
	})

LOOP:
	for {
		fmt.Println("---for循环一次---")
		select {
		case x, ok := <-c:
			fmt.Printf("%v, 通道读取到: x=%v,ok=%v\n", time.Now().Format(fmat), x, ok)
			time.Sleep(1 * time.Second)
		case <-timeout:
			fmt.Println("超时了")
			break LOOP
			/*default:
			fmt.Printf("%v,进入到default\n", time.Now().Format(fmat))
			// 下一次执行select的间隔取决于本次分支执行的时间长度，比如default执行了5秒，那就5秒后执行下一次的select
			time.Sleep(1000 * time.Millisecond)
			*/
		}
	}

	fmt.Println("方法结束")
}

// blockExample
// 结论:1.如果select里边只有一个case，而这个case被关闭了，则会死循环。
// 	   2.select里只有一个已经关闭的case,会deadlock
// 	   3.select会忽略设置为nil的chan
//**
func blockExample() {
	c := make(chan string, 1)
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("开始传输")
		c <- "你好"
		close(c)
	})

	for {
		fmt.Println("---for循环一次---")
		select {
		case x, ok := <-c:
			fmt.Printf("%v, 通道读取到: x=%v,ok=%v\n", time.Now().Format(fmat), x, ok)
			time.Sleep(1 * time.Second)

			if !ok {
				c = nil // 关闭后的chan设置为nil,select读取会阻塞
			}
		default:
			fmt.Printf("%v,进入到default\n", time.Now().Format(fmat))
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
