package main

func main() {
	/**
	堆内存布局


	1.arena
	Go语言的runtime将堆地址空间划分成了一个一个的arena，arena区域的起始地址被定义为常量arenaBaseOffset。
	在amd64架构的Linux环境下，每个arena的大小是64MB，起始地址也对齐到64MB，每个arena包含8192个page，所以每个page大小为8KB。

	2.采用了与tcmalloc内存分配器类似的算法。
	简单来讲就是：按照一组预置的大小规格把内存页划分成块，然后把不同规格的内存块放入对应的空闲链表中。

	Go 1.16 runtime包给出了67种预置的大小规格，最小8字节，最大32KB。
	所以在划分的整整齐齐的arena里，又会按需划分出不同的span，每个span包含一组连续的page，并且按照特定规格划分成了等大的内存块。

	3.arena,span,page
	arena, span, page和内存块组成了堆内存，而在堆内存之外，有一票用于管理堆内存的数据结构。

	例如，一个arena对应一个heapArena结构，一个span对应一个mspan结构。通过它们可以知道某个内存块是否已分配；已分配的内存用作指针还是标量；是否已被GC标记；是否等待清扫等信息。

	type heapArena struct {}
	type mspan struct {}

	type mheap struct {}
	type mcentral struct {}


	4.堆内存分配mallocgc
	这一次我们了解了负责分配内存的mallocgc()函数的主要逻辑，期间介绍了：
	（1）辅助GC；
	（2）三种内存分配策略；
	（3）从内存地址定位到heapArena和mspan的过程。

	maxSmallSize是32KB，maxTinySize等于16。也就是说：
	（1）小于16字节，而且是noscan类型的内存分配请求，会使用tiny allocator；
	（2）大于32KB的内存分配，包括noscan和scannable类型，都会采用大块内存分配器；
	（3）剩下的，大于等于16B且小于等于32KB的noscan类型；以及不大于32KB的scannable类型的分配请求，都会直接匹配预置的大小规格来分配。





































































	*/
}
