package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

/**
sync map原理:
1.空间换时间：通过冗余的两个数据结构(read、dirty)，实现加锁对性能的影响。
2.使用只读数据(read)，避免读写冲突。
3.动态调整，miss次数多了之后，将dirty数据迁移到read中。
4.double-checking。
5.延迟删除。 删除一个键值只是打标记，只有在迁移dirty数据的时候才清理删除的数据。
6.优先从read读取、更新、删除，因为对read的读取不需要锁。


总结：非常适合读多写少的情况
*/
func main() {

}

type Map struct {
	// 加锁作用，保护 dirty 字段
	mu sync.Mutex
	// 只读的数据，实际数据类型为 readOnly
	read atomic.Value
	// 最新写入的数据
	dirty map[interface{}]*entry
	// 计数器，每次需要读 dirty 则 +1
	misses int
}

type readOnly struct {
	// 内建 map
	m map[interface{}]*entry
	// 表示 dirty 里存在 read 里没有的 key，通过该字段决定是否加锁读 dirty
	amended bool
}

/*p有三种值：

nil: entry已被删除了，并且m.dirty为nil
expunged: entry已被删除了，并且m.dirty不为nil，而且这个entry不存在于m.dirty中
其它： entry是一个正常的值*/
type entry struct {
	p unsafe.Pointer // *interface{}
}
