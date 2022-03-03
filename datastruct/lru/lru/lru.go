package lru

/**
题目：https://leetcode-cn.com/problems/lru-cache/

LRU 缓存

请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int Capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
int get(int Key) 如果关键字 Key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int Key, int Value)如果关键字key 已经存在，则变更其数据值value ；如果不存在，则向缓存中插入该组key-Value 。如果插入操作导致关键字数量超过capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：

1 <= Capacity <= 3000
0 <= Key <= 10000
0 <= Value <= 105
最多调用 2 * 105 次 get 和 put
*/

// LRUCache Least Recently Used 最近最少使用，基于双向链表+哈希实现
type LRUCache struct {
	Size     int
	Capacity int
	Cache    map[int]*DoubleListNode
	// 固定从头部添加新元素，从尾部pop
	// 双向链表方便从尾部删除，单向链表不行
	DoubleList *DoubleList
}

type DoubleList struct {
	Head, Tail *DoubleListNode
}

type DoubleListNode struct {
	Key, Value int
	Prev, Next *DoubleListNode
}

func initDoubleListNode(key, value int) *DoubleListNode {
	return &DoubleListNode{
		Key:   key,
		Value: value,
	}
}

func initDoubleList() *DoubleList {
	d := &DoubleList{
		Head: initDoubleListNode(0, 0),
		Tail: initDoubleListNode(0, 0),
	}

	// 初始化的时候有2个结点，头结点，尾结点互相指
	d.Head.Next = d.Tail
	d.Tail.Prev = d.Head

	return d
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Cache:      map[int]*DoubleListNode{},
		Capacity:   capacity,
		DoubleList: initDoubleList(),
	}

	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}

	node := this.Cache[key]
	this.moveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}

	if _, ok := this.Cache[key]; !ok {
		if this.Size+1 > this.Capacity {
			removed := this.removeTail()
			delete(this.Cache, removed.Key)
		} else {
			this.Size++
		}

		node := initDoubleListNode(key, value)
		this.Cache[key] = node
		this.addToHead(node)
	} else {
		node := this.Cache[key]
		node.Value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DoubleListNode) {
	// 不是替换头结点，而是添加到头结点后边一个元素
	node.Prev = this.DoubleList.Head
	node.Next = this.DoubleList.Head.Next
	this.DoubleList.Head.Next.Prev = node
	this.DoubleList.Head.Next = node
}

func (this *LRUCache) removeNode(node *DoubleListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	node.Next = nil
	node.Prev = nil
}

func (this *LRUCache) moveToHead(node *DoubleListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 超出容量的时候需要移除元素
func (this *LRUCache) removeTail() *DoubleListNode {
	node := this.DoubleList.Tail.Prev
	this.removeNode(node)
	return node
}
