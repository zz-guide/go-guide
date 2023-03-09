########channel底层源码分析

不要通过共享内存的方式进行通信，而是应该通过通信的方式共享内存

虽然我们在 Go 语言中也能使用共享内存加互斥锁进行通信，
但是 Go 语言提供了一种不同的并发模型，即通信顺序进程（Communicating sequential processes，CSP）1。
Goroutine 和 Channel 分别对应 CSP 中的实体和传递信息的媒介，Goroutine 之间会通过 Channel 传递数据。

1.Channel 收发操作均遵循了先进先出的设计,带缓冲区和不带缓冲区的 Channel 都会遵循先入先出发送和接收数据。
2.乐观锁
锁是一种常见的并发控制技术，我们一般会将锁分成乐观锁和悲观锁，即乐观并发控制和悲观并发控制，
无锁（lock-free）队列更准确的描述是使用乐观并发控制的队列。
乐观并发控制也叫乐观锁，很多人都会误以为乐观锁是与悲观锁差不多，
然而它并不是真正的锁，只是一种并发控制的思想。

乐观并发控制本质上是基于验证的协议，我们使用原子指令 CAS（compare-and-swap 或者 compare-and-set）在多线程中同步数据，无锁队列的实现也依赖这一原子指令

3.runtime.hchan结构
该结构体中包含了用于保护成员变量的互斥锁，从某种程度上说，
Channel 是一个用于同步和通信的有锁队列，使用互斥锁解决程序中可能存在的线程竞争问题是很常见的，我们能很容易地实现有锁队列。

而 Go 语言社区也在 2014 年提出了无锁 Channel 的实现方案，该方案将 Channel 分成了以下三种类型8：

同步 Channel — 不需要缓冲区，发送方会直接将数据交给（Handoff）接收方；
异步 Channel — 基于环形缓存的传统生产者消费者模型；
chan struct{} 类型的异步 Channel — struct{} 类型不占用内存空间，不需要实现缓冲区和直接发送（Handoff）的语义；
这个提案的目的也不是实现完全无锁的队列，只是在一些关键路径上通过无锁提升 Channel 的性能。社区中已经有无锁 Channel 的实现9，但是在实际的基准测试中，无锁队列在多核测试中的表现还需要进一步的改进10。

因为目前通过 CAS 实现11的无锁 Channel 没有提供先进先出的特性，所以该提案暂时也被搁浅了12。

type hchan struct {
    qcount   uint           // total data in the queue
    dataqsiz uint           // size of the circular queue
    buf      unsafe.Pointer // points to an array of dataqsiz elements
    elemsize uint16
    closed   uint32
    elemtype *_type // element type
    sendx    uint   // send index
    recvx    uint   // receive index
    recvq    waitq  // list of recv waiters
    sendq    waitq  // list of send waiters
    lock mutex
    // lock protects all fields in hchan, as well as several
    // fields in sudogs blocked on this channel.
    //
    // Do not change another G's status while holding this lock
    // (in particular, do not ready a G), as this can deadlock
    // with stack shrinking.
   
}

type waitq struct {
    first *sudog
    last  *sudog
}

buf环形缓冲区
channel类型，大小，数量，收发队列，当前读写位置，是否关闭，互斥锁

4.创建
make关键字创建
无缓冲

有缓冲


5.发送数据

ch <- i 语句
发送消息的时候会先加锁，防止多个协程并发修改数据，如果已经关闭直接panic

三部分：
当存在等待的接收者时，通过 runtime.send 直接将数据发送给阻塞的接收者；
当缓冲区存在空余空间时，将发送的数据写入 Channel 的缓冲区；
当不存在缓冲区或者缓冲区已满时，等待其他 Goroutine 从 Channel 接收数据；

直接发送：
    1.数据直接拷贝
    2.唤醒等待的g，等待被下一次调度


6.接收数据
两种方式接收
i <- ch
i, ok <- ch

当存在等待的发送者时，通过 runtime.recv 从阻塞的发送者或者缓冲区中获取数据；
当缓冲区存在数据时，从 Channel 的缓冲区中接收数据；
当缓冲区中不存在数据时，等待其他 Goroutine 向 Channel 发送数据；

直接接收
从发送队列获取数据
从缓冲区中接接收数据

7.关闭close
1.当 Channel 是一个空指针或者已经被关闭时，Go 语言运行时都会直接崩溃并抛出异常：
2.该函数在最后会为所有被阻塞的 Goroutine 调用 runtime.goready 触发调度。


