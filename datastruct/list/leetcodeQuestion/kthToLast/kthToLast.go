package main

import (
	. "go-guide/datastruct/list/node"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/

返回倒数第 k 个节点

实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k保证是有效的。


*/

func main() {
	list := MakeListNode([]int{1, 2, 3, 4})
	k := 2
	log.Println("返回倒数第k个节点:", kthToLast(list, k))
}

// kthToLast 快慢指针，时间复杂度O(n)，空间复杂度O(1)
func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	// fast先走k步，到k+1节点
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	// 判断边界情况
	if fast == nil {
		return head.Val
	}

	// 两指针同时移动1
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
