# Go的内存分配

# 参考:https://cloud.tencent.com/developer/article/2179344

栈内存分配
小于32KB的栈内存
来源优先级1：线程缓存mcache
来源优先级2：全局缓存stackpool
来源优先级3：逻辑处理器结构p.pagecache
来源优先级4：堆mheap
大于等于32KB的栈内存
来源优先级1：全局缓存stackLarge
来源优先级2：逻辑处理器结构p.pagecache
来源优先级3：堆mheap
堆内存分配
微对象 0 < Micro Object < 16B
来源优先级1：线程缓存mcache.tiny
来源优先级2：线程缓存mcache.alloc
小对象 16B =< Small Object <= 32KB
来源优先级1：线程缓存mcache.alloc
来源优先级2：中央缓存mcentral
来源优先级3：逻辑处理器结构p.pagecache
来源优先级4：堆mheap
大对象 32KB < Large Object
来源优先级1：逻辑处理器结构p.pagecache
来源优先级2：堆mheap
「栈内存」也来源于堆mheap