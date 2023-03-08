########定时器底层源码分析


1.Mutex
Go 语言的 sync.Mutex 由两个字段 state 和 sema 组成。其中 state 表示当前互斥锁的状态，而 sema 是用于控制锁状态的信号量。
type Mutex struct {
    state int32
    sema  uint32
}
上述两个加起来只占 8 字节空间的结构体表示了 Go 语言中的互斥锁。
互斥锁的状态比较复杂，如下图所示，最低三位分别表示 mutexLocked、mutexWoken 和 mutexStarving，剩下的位置用来表示当前有多少个 Goroutine 在等待互斥锁的释放：

在默认情况下，互斥锁的所有状态位都是 0，int32 中的不同位分别表示了不同的状态：

mutexLocked — 表示互斥锁的锁定状态；
mutexWoken — 表示从正常模式被从唤醒；
mutexStarving — 当前的互斥锁进入饥饿状态；
waitersCount — 当前互斥锁上等待的 Goroutine 个数；



2.正常模式和饥饿模式
在正常模式下，锁的等待者会按照先进先出的顺序获取锁。但是刚被唤起的 Goroutine 与新创建的 Goroutine 竞争时，
大概率会获取不到锁，为了减少这种情况的出现，一旦 Goroutine 超过 1ms 没有获取到锁，
它就会将当前互斥锁切换饥饿模式，防止部分 Goroutine 被『饿死』。

引入的目的是保证互斥锁的公平性。

在饥饿模式中，互斥锁会直接交给等待队列最前面的 Goroutine。新的 Goroutine 在该状态下不能获取锁、也不会进入自旋状态，它们只会在队列的末尾等待。如果一个 Goroutine 获得了互斥锁并且它在队列的末尾或者它等待的时间少于 1ms，那么当前的互斥锁就会切换回正常模式。
与饥饿模式相比，正常模式下的互斥锁能够提供更好地性能，饥饿模式的能避免 Goroutine 由于陷入等待无法获取锁而造成的高尾延时。


如果互斥锁的状态不是 0 时就会调用 sync.Mutex.lockSlow 尝试通过自旋（Spinnig）等方式等待锁的释放，该方法的主体是一个非常大 for 循环，这里将它分成几个部分介绍获取锁的过程：

判断当前 Goroutine 能否进入自旋；
通过自旋等待互斥锁的释放；
计算互斥锁的最新状态；
更新互斥锁的状态并获取锁；
我们先来介绍互斥锁是如何判断当前 Goroutine 能否进入自旋等互斥锁的释放：

自旋是一种多线程同步机制，当前的进程在进入自旋的过程中会一直保持 CPU 的占用，持续检查某个条件是否为真。在多核的 CPU 上，自旋可以避免 Goroutine 的切换，使用恰当会对性能带来很大的增益，但是使用的不恰当就会拖慢整个程序，所以 Goroutine 进入自旋的条件非常苛刻：

互斥锁只有在普通模式才能进入自旋；
runtime.sync_runtime_canSpin 需要返回 true：
运行在多 CPU 的机器上；
当前 Goroutine 为了获取该锁进入自旋的次数小于四次；
当前机器上至少存在一个正在运行的处理器 P 并且处理的运行队列为空；
一旦当前 Goroutine 能够进入自旋就会调用runtime.sync_runtime_doSpin 和 runtime.procyield 并执行 30 次的 PAUSE 指令，该指令只会占用 CPU 并消耗 CPU 时间：

互斥锁的解锁过程 sync.Mutex.Unlock 与加锁过程相比就很简单，该过程会先使用 sync/atomic.AddInt32 函数快速解锁，这时会发生下面的两种情况：

如果该函数返回的新状态等于 0，当前 Goroutine 就成功解锁了互斥锁；
如果该函数返回的新状态不等于 0，这段代码会调用 sync.Mutex.unlockSlow 开始慢速解锁：


sync.Mutex.unlockSlow 会先校验锁状态的合法性 — 如果当前互斥锁已经被解锁过了会直接抛出异常 “sync: unlock of unlocked mutex” 中止当前程序。

在正常情况下会根据当前互斥锁的状态，分别处理正常模式和饥饿模式下的互斥锁：

在正常模式下，上述代码会使用如下所示的处理过程：
如果互斥锁不存在等待者或者互斥锁的 mutexLocked、mutexStarving、mutexWoken 状态不都为 0，那么当前方法可以直接返回，不需要唤醒其他等待者；
如果互斥锁存在等待者，会通过 sync.runtime_Semrelease 唤醒等待者并移交锁的所有权；
在饥饿模式下，上述代码会直接调用 sync.runtime_Semrelease 将当前锁交给下一个正在尝试获取锁的等待者，等待者被唤醒后会得到锁，在这时互斥锁还不会退出饥饿状态；




3.读写锁
获取写锁时会先阻塞写锁的获取，后阻塞读锁的获取，这种策略能够保证读操作不会被连续的写操作『饿死』。


