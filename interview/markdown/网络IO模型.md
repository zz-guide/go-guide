########网络IO模型


1.网络轮询器不仅用于监控网络 I/O，还能用于监控文件的 I/O，它利用了操作系统提供的 I/O 多路复用模型来提升 I/O 设备的利用率以及程序的性能。

I/O 模型 #
操作系统中包含阻塞 I/O、非阻塞 I/O、信号驱动 I/O 与异步 I/O 以及 I/O 多路复用五种 I/O 模型。我们在本节中会介绍上述五种模型中的三种：

阻塞 I/O 模型；
非阻塞 I/O 模型；
I/O 多路复用模型；
在 Unix 和类 Unix 操作系统中，文件描述符（File descriptor，FD）是用于访问文件或者其他 I/O 资源的抽象句柄，例如：管道或者网络套接字1。而不同的 I/O 模型会使用不同的方式操作文件描述符。


2.多模块
Go 语言在网络轮询器中使用 I/O 多路复用模型处理 I/O 操作，但是他没有选择最常见的系统调用 select2。虽然 select 也可以提供 I/O 多路复用的能力，但是使用它有比较多的限制：

监听能力有限 — 最多只能监听 1024 个文件描述符；
内存拷贝开销大 — 需要维护一个较大的数据结构存储文件描述符，该结构需要拷贝到内核中；
时间复杂度 𝑂(𝑛) — 返回准备就绪的事件个数后，需要遍历所有的文件描述符；

Go 语言为了提高在不同操作系统上的 I/O 操作性能，使用平台特定的函数实现了多个版本的网络轮询模块：

src/runtime/netpoll_epoll.go
src/runtime/netpoll_kqueue.go
src/runtime/netpoll_solaris.go
src/runtime/netpoll_windows.go
src/runtime/netpoll_aix.go
src/runtime/netpoll_fake.go

这些模块在不同平台上实现了相同的功能，构成了一个常见的树形结构。编译器在编译 Go 语言程序时，会根据目标平台选择树中特定的分支进行编译：

如果目标平台是 Linux，那么就会根据文件中的 // +build linux 编译指令选择 src/runtime/netpoll_epoll.go 并使用 epoll 函数处理用户的 I/O 操作。

接口 #
epoll、kqueue、solaries 等多路复用模块都要实现以下五个函数，这五个函数构成一个虚拟的接口：

func netpollinit()
func netpollopen(fd uintptr, pd *pollDesc) int32
func netpoll(delta int64) gList
func netpollBreak()
func netpollIsPollDescriptor(fd uintptr) bool

上述函数在网络轮询器中分别扮演了不同的作用：

runtime.netpollinit — 初始化网络轮询器，通过 sync.Once 和 netpollInited 变量保证函数只会调用一次；
runtime.netpollopen — 监听文件描述符上的边缘触发事件，创建事件并加入监听；
runtime.netpoll — 轮询网络并返回一组已经准备就绪的 Goroutine，传入的参数会决定它的行为3；
如果参数小于 0，无限期等待文件描述符就绪；
如果参数等于 0，非阻塞地轮询网络；
如果参数大于 0，阻塞特定时间轮询网络；
runtime.netpollBreak — 唤醒网络轮询器，例如：计时器向前修改时间时会通过该函数中断网络轮询器4；
runtime.netpollIsPollDescriptor — 判断文件描述符是否被轮询器使用；
我们在这里只需要了解多路复用模块中的几个函数，本节的后半部分会详细分析它们的实现原理。


3.数据结构

操作系统中 I/O 多路复用函数会监控文件描述符的可读或者可写，而 Go 语言网络轮询器会监听 runtime.pollDesc 结构体的状态，它会封装操作系统的文件描述符：

type pollDesc struct {
    link *pollDesc
	lock    mutex
	fd      uintptr
	...
	rseq    uintptr
	rg      uintptr
	rt      timer
	rd      int64
	wseq    uintptr
	wg      uintptr
	wt      timer
	wd      int64
}

该结构体中包含用于监控可读和可写状态的变量，我们按照功能将它们分成以下四组：

rseq 和 wseq — 表示文件描述符被重用或者计时器被重置5；
rg 和 wg — 表示二进制的信号量，可能为 pdReady、pdWait、等待文件描述符可读或者可写的 Goroutine 以及 nil；
rd 和 wd — 等待文件描述符可读或者可写的截止日期；
rt 和 wt — 用于等待文件描述符的计时器；
除了上述八个变量之外，该结构体中还保存了用于保护数据的互斥锁、文件描述符。runtime.pollDesc 结构体会使用 link 字段串联成链表存储在 runtime.pollCache 中：


type pollCache struct {
    lock  mutex
    first *pollDesc
}


4.轮询缓存链表


5.多路复用
网络轮询器实际上是对 I/O 多路复用技术的封装，本节将通过以下的三个过程分析网络轮询器的实现原理：
网络轮询器的初始化；
如何向网络轮询器加入待监控的任务；
如何从网络轮询器获取触发的事件；


6.事件循环
我们将从以下的两个部分介绍事件循环的实现原理：
Goroutine 让出线程并等待读写事件；
多路复用等待读写事件的发生并返回；

等待事件
轮询等待
截止日期

网络轮询器和计时器的关系非常紧密，这不仅仅是因为网络轮询器负责计时器的唤醒，还因为文件和网络 I/O 的截止日期也由网络轮询器负责处理。截止日期在 I/O 操作中，尤其是网络调用中很关键，网络请求存在很高的不确定因素，我们需要设置一个截止日期保证程序的正常运行，这时需要用到网络轮询器中的 runtime.poll_runtime_pollSetDeadline：




