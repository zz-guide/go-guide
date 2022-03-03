package lfu

/**
题目：https://leetcode-cn.com/problems/lfu-cache/
LFU 缓存

请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量capacity 初始化对象
int get(int Key)- 如果键key 存在于缓存中，则获取键的值，否则返回 -1 。
void put(int Key, int value)- 如果键key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入：
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

解释：
// cnt(x) = 键 x 的使用计数
// Cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // Cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // Cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // 返回 1
                 // Cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
                 // Cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // Cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
                 // Cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // Cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // 返回 4
                 // Cache=[3,4], cnt(4)=2, cnt(3)=3

提示：

0 <= capacity<= 104
0 <= Key <= 105
0 <= value <= 109
最多调用 2 * 105 次 get 和 put 方法

*/

type DoubleList struct {
	Head, Tail *DoubleListNode
}

type DoubleListNode struct {
	Prev, Next       *DoubleListNode
	Key, Value, Freq int
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

func (this *DoubleList) addToHead(node *DoubleListNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head

	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *DoubleList) Remove(node *DoubleListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	node.Next = nil
	node.Prev = nil
}

func (this *DoubleList) RemoveLast() *DoubleListNode {
	if this.IsEmpty() {
		return nil
	}

	last := this.Tail.Prev
	this.Remove(last)

	return last
}

func (this *DoubleList) IsEmpty() bool {
	return this.Head.Next == this.Tail
}

type LFUCache struct {
	Cache                   map[int]*DoubleListNode
	Freq                    map[int]*DoubleList
	Capacity, Size, MinFreq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Cache:    make(map[int]*DoubleListNode),
		Freq:     make(map[int]*DoubleList),
		Capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.IncFreq(node)
		return node.Value
	}
	return -1
}

func (this *LFUCache) IncFreq(node *DoubleListNode) {
	_freq := node.Freq
	this.Freq[_freq].Remove(node)
	if this.MinFreq == _freq && this.Freq[_freq].IsEmpty() {
		this.MinFreq++
		delete(this.Freq, _freq)
	}

	node.Freq++
	if this.Freq[node.Freq] == nil {
		this.Freq[node.Freq] = initDoubleList()
	}
	this.Freq[node.Freq].addToHead(node)
}

func (this *LFUCache) Put(key, value int) {
	if this.Capacity == 0 {
		return
	}

	if node, ok := this.Cache[key]; ok {
		node.Value = value
		this.IncFreq(node)
	} else {
		if this.Size+1 > this.Capacity {
			node := this.Freq[this.MinFreq].RemoveLast()
			delete(this.Cache, node.Key)
			this.Size--
		}

		x := &DoubleListNode{Key: key, Value: value, Freq: 1}
		this.Cache[key] = x
		if this.Freq[1] == nil {
			this.Freq[1] = initDoubleList()
		}
		this.Freq[1].addToHead(x)
		this.MinFreq = 1
		this.Size++
	}
}
