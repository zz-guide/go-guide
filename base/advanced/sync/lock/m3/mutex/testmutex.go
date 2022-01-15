package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {

}

func coment() {
	/**
	结论一：sync.Mutex表示互斥锁，同时只有一个协程可以持有资源，只有释放了以后其他协程才可以使用
	*/
}

/**
sync.Mutex，sync.RMutex，sync.Once，sync.Cond，sync.Waitgroup,sync.Pool
*/
func TestMutex() {
	var a = 0
	lock := sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}

	// 等待 1s 结束主程序,确保所有协程执行完
	time.Sleep(time.Second * 1)
}

/**
结论一：在首次使用后不要复制该互斥锁。对一个未锁定的互斥锁解锁将会产生运行时错误。
结论二：一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）
结论三：互斥锁不释放会造成死锁，fatal error: all goroutines are asleep - deadlock!
*/
func TestMutex2() {
	ch := make(chan string, 2)

	var l sync.Mutex //互斥锁
	go func() {
		l.Lock() //如果直接释放锁的话，fatal error: sync: unlock of unlocked mutex
		defer l.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你们去抢吧")
		ch <- "nihao"
	}()

	go func() {
		fmt.Println("goroutine2: 阻塞中，等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 哈哈，我锁定了")
		ch <- "shijie"
	}()

	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
}

/**
读锁定（RLock），对读操作进行锁定
读解锁（RUnlock），对读锁定进行解锁
写锁定（Lock），对写操作进行锁定
写解锁（Unlock），对写锁定进行解锁
在首次使用之后，不要复制该读写锁。不要混用锁定和解锁，
如：Lock 和 RUnlock、RLock 和 Unlock。
因为对未读锁定的读写锁进行读解锁或对未写锁定的读写锁进行写解锁将会引起运行时错误。

也就是说，当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；
当有一个 goroutine 获得读锁定，其它读锁定任然可以继续；当有一个或任意多个读锁定，
写锁定将等待所有读锁定解锁之后才能够进行写锁定。所以说这里的读锁定（RLock）目的其实是告诉写锁定：
有很多人正在读取数据，你给我站一边去，等它们读（读解锁）完你再来写（写锁定）。
*/

var count int

var rw sync.RWMutex

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

/**
rand.Seed(time.Now().UnixNano())		//若不设定随机数种子，每次取的都是一样的值
fmt.Println(rand.Intn(1000))
*/

func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}

func TestRWMetux() {
	ch := make(chan struct{}, 10)

	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
}

/**
WaitGroup 用于等待一组 goroutine 结束，用法很简单。它有三个方法

func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()

Add 用来添加 goroutine 的个数。Done 执行一次数量减 1。Wait 用来等待结束

*/

func TestWaitGroup() {

	seconds := []int{1, 2, 3}

	var wg sync.WaitGroup

	for i, s := range seconds {
		// 计数加 1
		wg.Add(2)
		go func(i, s int) {
			time.Sleep(time.Second * 2)
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine%d 结束\n", i)
		}(i, s)

		go func(i, s int) {
			time.Sleep(time.Second * 1)
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine%d 结束\n", i)
		}(i, s)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
}

/**
使用 sync.Once 对象可以使得函数多次调用只执行一次。用 done 来记录执行次数，用 m 来保证保证仅被执行一次。只有一个 Do 方法，调用执行。
*/
func TestOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			//onceBody()
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

/**
Pool 临时对象池
sync.Pool 可以作为临时对象的保存和复用的集合。
*/

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	// 获取临时对象，没有的话会自动创建
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	// 将临时对象放回到 Pool 中
	bufPool.Put(b)
}

func TestPool() {
	Log(os.Stdout, "path", "/search?q=flowers")
}

/**
Cond 条件变量,Cond 实现一个条件变量，即等待或宣布事件发生的 goroutines 的会合点。
Wait 方法、Signal 方法和 Broadcast 方法。它们分别代表了等待通知、单发通知和广播通知的操作。
*/
func TestCond() {
	count := 2

	ch := make(chan struct{}, 5)

	var l sync.Mutex
	cond := sync.NewCond(&l) //给某一个互斥锁增加一个cond

	for i := 0; i < 5; i++ {
		go func(i int) {
			// 争抢互斥锁的锁定
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()

			// 条件是否达成
			for count > i {
				cond.Wait()
				fmt.Printf("收到一个通知 goroutine%d\n", i)
			}

			fmt.Printf("goroutine%d 执行结束\n", i)
		}(i)
	}

	// 确保所有 goroutine 启动完成
	time.Sleep(time.Millisecond * 20)
	// 锁定一下，我要改变 count 的值
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	count -= 2
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	for i := 0; i < 5; i++ {
		<-ch
	}
}
