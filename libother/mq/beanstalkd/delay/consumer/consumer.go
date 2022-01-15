package main

import (
	"fmt"
	"github.com/beanstalkd/go-beanstalk"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**

beanstalkd 是一个轻量级消息中间件，其主要特性：
基于管道  (tube) 和任务 (job) 的工作队列 (work-queue)：
管道（tube），tube类似于消息主题（topic），在一个beanstalkd中可以支持多个tube，每个tube都有自己的producer和consumer；
任务（job），beanstalkd用job代替了message的概念，与消息不同，job有一系列状态：
内部实现采用了 libevent, 服务器-客户端之间用类似 memcached 的轻量级通讯协议，具有有很高的性能。
尽管是内存队列，beanstalkd 提供了 binlog 机制，当重启 beanstalkd 时，当前任务状态能够从纪录的本地 binlog 中恢复。
优先级（priority）：job可以有0~2^32个优先级，0代表最高优先级，beanstalkd使用最大最小堆处理job的优先级排序，因此reserve命令的时间复杂度是O(logN)；
延时（delay），有两种方式可以执行延时任务：producer发布任务时指定延时；或者当任务处理完毕后, consumer再次将任务放入队列延时执行 (RELEASE with <delay>)；
超时重发（time-to-run），Beanstalkd 把job返回给consumer以后：consumer必须在预设的 TTR (time-to-run) 时间内发送 delete / release/ bury 改变任务状态；否则 Beanstalkd 会认为消息处理失败，然后把job交给另外的消费者节点执行。如果consumer预计在 TTR (time-to-run) 时间内无法完成任务, 也可以发送 touch 命令, 它的作用是让 Beanstalkd 从系统时间重新计算 TTR ；
任务预留（buried），如果job因为某些原因无法执行, consumer可以把任务置为 buried 状态让 Beanstalkd 保留这些任务。管理员可以通过 peek buried 命令查询被保留的任务，并且进行人工干预。简单的, kick <n> 能够一次性把 n 条被保留的任务踢回队列。

job的状态
READY，需要立即处理的任务，当延时 (DELAYED) 任务到期后会自动成为当前任务；
DELAYED，延迟执行的任务, 当消费者处理任务后，可以用将消息再次放回 DELAYED 队列延迟执行；
RESERVED，已经被消费者获取, 正在执行的任务，Beanstalkd 负责检查任务是否在 TTR(time-to-run) 内完成；
BURIED，保留的任务: 任务不会被执行，也不会消失，除非有人把它 "踢" 回队列；
DELETED，消息被彻底删除。Beanstalkd 不再维持这些消息。

问题:
DEADLINE_SOON是什么意思?
DEADLINE_SOON是一个reserve命令的响应，它表明一个“reserved”状态的job的最后期限（deadline）马上要到期（目前的安全边际大约是1秒）。

如果你执行reserve命令时，频繁地收到DEADLINE_SOON错误，你可能应该考虑对你的job增加TTR，因为它表示你没有安时完成你的job。这也可能是在完成了你的job，却没有删除它们。

*/
var (
	tube       = "tube1"
	reserveTtl = 5 * time.Second // ttl时间内没有获取job都会立即返回,-1表示没有消息一直阻塞
)

func main() {
	log.Println("启动消费者")
	ReceiveMessage()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("main主动退出")
}

func ReceiveMessage() {
	conn, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		log.Println("connect 出错:", err)
		return
	}

	tubeSet := beanstalk.NewTubeSet(conn, tube)
	go func() {
		// 典型的pull模式
		for {
			fmt.Println("开始拉取信息...")
			// 一个一个job取出来
			id, body, err := tubeSet.Reserve(reserveTtl)
			log.Printf("id: %d, body: %s, err: %s \n", id, string(body), err)
			if err != nil {
				// the error can only be beanstalk.NameError or beanstalk.ConnError
				switch cErr := err.(type) {
				case beanstalk.ConnError:
					switch cErr.Err {
					case beanstalk.ErrTimeout:
						// timeout error on timeout, just continue the loop
						continue
					case beanstalk.ErrBadChar, beanstalk.ErrBadFormat, beanstalk.ErrBuried, beanstalk.ErrDeadline,
						beanstalk.ErrDraining, beanstalk.ErrEmpty, beanstalk.ErrInternal, beanstalk.ErrJobTooBig,
						beanstalk.ErrNoCRLF, beanstalk.ErrNotFound, beanstalk.ErrNotIgnored, beanstalk.ErrTooLong:
						// won't reset
						log.Println(err)
					default:
						// beanstalk.ErrOOM, beanstalk.ErrUnknown and other errors
						log.Println(err)
						time.Sleep(time.Second)
					}
				default:
					log.Println(err)
				}

				continue
			}

			if len(body) > 0 {
				// 消费
				log.Println("接收消息:", string(body))
				conn.Delete(id) // 删除消息
				//conn.Touch(id)                      // 没有处理完毕，重新设置ttl
				//conn.Release(id, 0, 30*time.Second) // 重新放回队列,30秒之后才可以被再次消费
				//conn.Bury(id, 0)                    // bury命令将job放到一个特殊的FIFO队列中，之后不能被reserve命令获取，但可以用kick命令扔回工作队列中，之后就能被消费了：
				time.Sleep(time.Second * 1)
				continue
			}

		}
	}()
}
