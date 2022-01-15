package main

/**
题目：https://leetcode-cn.com/problems/design-a-stack-with-increment-operation/

请你设计一个支持下述操作的栈。

实现自定义栈类 CustomStack ：

CustomStack(int maxSize)：用 maxSize 初始化对象，maxSize 是栈中最多能容纳的元素数量，栈在增长到 maxSize 之后则不支持 push 操作。
void push(int x)：如果栈还未增长到 maxSize ，就将 x 添加到栈顶。
int pop()：弹出栈顶元素，并返回栈顶的值，或栈为空时返回 -1 。
void inc(int k, int val)：栈底的 k 个元素的值都增加 val 。如果栈中元素总数小于 k ，则栈中的所有元素都增加 val 。

提示：

1 <= maxSize <= 1000
1 <= x <= 1000
1 <= k <= 1000
0 <= val <= 100
每种方法 increment，push 以及 pop 分别最多调用 1000 次

*/
func main() {

}

type CustomStack struct {
	stack []int
	c     int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		c:     0,
		stack: make([]int, maxSize, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if this.c < len(this.stack) {
		this.stack[this.c] = x
		this.c++
	}
}

func (this *CustomStack) Pop() int {
	if this.c != 0 {
		this.c--
		return this.stack[this.c]
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	if k > this.c {
		for i := 0; i < this.c; i++ {
			this.stack[i] += val
		}
	} else {
		for i := 0; i < k; i++ {
			this.stack[i] += val
		}
	}
}
