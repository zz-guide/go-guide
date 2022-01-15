package main

import (
	"fmt"
	"math"
	"unsafe"
)

const bucketCntBits = 3
const bucketCnt = 1 << bucketCntBits

type hmap struct {
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	nextOverflow *bmap
}

type bmap struct {
	tophash [bucketCnt]uint8
}

//type Map struct {
//	Key  *Type // Key type
//	Elem *Type // Val (elem) type
//
//	Bucket *Type // internal struct type representing a hash bucket
//	Hmap   *Type // internal struct type representing the Hmap (map header object)
//	Hiter  *Type // internal struct type representing hash iterator state
//}

// 编译期间变成了这种结构
// tophash包含此桶中每个键的哈希值最高字节（高8位）信息（也就是前面所述的high-order bits）。
// 如果tophash[0] < minTopHash，tophash[0]则代表桶的搬迁（evacuation）状态。
//type bmap1 struct {
//	topbits  [8]uint8
//	keys     [8]keytype
//	values   [8]valuetype
//	pad      uintptr
//	overflow uintptr
//}

/**
结论：
1.涉及到的结构体有hmap,bmap,mapextra
2.bmap就是最终存数据的桶，每个桶装固定的8对kv，排列的话先是8个key在高位，地位是8个value，依次排列
3.多余8个时,会产生overflow桶
4.hash函数：在程序启动时，会检测 cpu 是否支持 aes，如果支持，则使用 aes hash，否则使用 memhash
5.触发扩容的装载因子为13/2=6.5
6.一个桶中最多能装载的键值对（key-value）的个数为8
7.键和值超过128个字节，就会被转换为指针
8.对于不指定初始化大小，和初始化值hint<=8（bucketCnt）时，go会调用makemap_small函数（源码位置src/runtime/map.go），并直接从堆上进行分配。
9.正常桶和溢出桶在内存中的存储空间是连续的，只是被 hmap 中的不同字段引用而已。
10.当hint>8时，则调用makemap函数
11.对于哈希算法的选择，程序会根据当前架构判断是否支持AES，如果支持就使用AES hash，其实现代码位于src/runtime/asm_{386,amd64,arm64}.s中；若不支持，其hash算法则根据xxhash算法（https://code.google.com/p/xxhash/）和cityhash算法（https://code.google.com/p/cityhash/）启发而来，代码分别对应于32位（src/runtime/hash32.go）和64位机器（src/runtime/hash32.go）中，对这部
不同的系统架构使用不同的hash函数
12.哈希值低位（low-order bits）用于选择桶，哈希值高位（high-order bits）用于在一个独立的桶中区别出键
13.golang采用位运算计算桶的位置，不是取余
14.go采用拉链法解决hash冲突，本桶无位置时，存到溢出桶，没有就添加溢出桶存
15.hash code高位用来查找元素
16.遍历是无序的，因为有fastRand()
17.元素个数 >= 桶（bucket）总数 * 6.5时，判断需要挂在溢出桶
18.判断溢出桶是否太多，当桶总数 < 2 ^ 15 时，如果溢出桶总数 >= 桶总数，则认为溢出桶过多。当桶总数 >= 2 ^ 15 时，直接与 2 ^ 15 比较，当溢出桶总数 >= 2 ^ 15 时，即认为溢出桶太多了。
19.不断的增删会导致溢出桶增多，但是元素又很少

20.增量扩容：翻倍
21.等量扩容，实际buckets数量不变，重新搬迁。go每次构造map都会重新生成hash种子，防止退化为链表
22.只有assign和delete的时候才会真正的搬迁，防止瞬时抖动
23.go采用的是渐进式hash，把旧的buckets分配到新的2个buckets中
24.定义了2^B个桶
*/
func main() {
	//F1()
	//F2()
	F3()
	//fmt.Printf("%b  %b %b %b\n", 1, 1<<1, 1<<2, 1<<3)
}

// F1
// 1.遍历无序
// 2.类似slice，会影响外部map,底层公用
// 3.并发读写不安全
// 4.不要使用math.NaN()作为key，因为math.NaN() 不等于自身，每次hash结果都不一样
// 5.hash表(哈希桶+数组)
///**
func F1() {
	// 结论：对于map类型，只有长度，没有容量
	m := make(map[string]string, 0)
	m["a"] = "a"
	fmt.Println("m:", len(m))
	fmt.Printf("m:%#v\n", m)

	fmt.Println(math.NaN())

	c := m
	c["a"] = "c"
	fmt.Println("c:", len(c), c)
	fmt.Println("m:", len(m), m)
}

// F2
// 缩容的条件：条件为溢出桶（noverflow）的数量 >= 32768（1<<15）
// 这种方式其实是存在运行隐患的，也就是导致在删除元素时，并不会释放内存，使得分配的总内存不断增加。如果一个不小心，拿 map 来做大 key/value 的存储，也不注意管理，很容易就内存爆了。
//
//也就是 Go 语言的 map 目前实现的是 ”伪缩容“，仅针对溢出桶过多的情况。若是触发缩容，hash 数组的占用的内存大小不变。
//
//若要实现 ”真缩容“，Go Contributor @josharian 表示目前唯一可用的解决方法是：创建一个新的 map 并从旧的 map 中复制元素。
///**
func F2() {
	// 换一个map可以解决内存问题，把大key删掉
	oldMap := make(map[int]int, 100000)
	fmt.Println("oldMap:", len(oldMap))
	newMap := make(map[int]int, len(oldMap))
	fmt.Println("newMap:", len(newMap))
	//for k, v := range oldMap {
	//	newMap[k] = v
	//}

	//oldMap = newMap
	//fmt.Println("newMap:", len(newMap))
	//fmt.Println("oldMap:", len(oldMap))
}

func F3() {
	m := map[string]string{}
	m["a"] = "s6"

	v, ok := m["a"]
	if ok {
		fmt.Println("v:", v)
	} else {
		fmt.Println("没找到")
	}

}
