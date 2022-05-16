package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/

K个一组翻转链表

给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。

k是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

提示：

列表中节点的数量在范围 sz 内
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz


思路：
	1.准备一个函数用来反转给定的链表
	2.数k个,反转

*/
func main() {
	list1 := MakeListNode([]int{1, 2, 3, 4, 5})
	k := 2
	PrintListNode(reverseKGroup(list1, k))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// myReverse 专门用来反转一段链表
	var myReverse func(head, tail *ListNode) (*ListNode, *ListNode)
	myReverse = func(head, tail *ListNode) (*ListNode, *ListNode) {
		// nil也不影响
		prev := tail.Next
		p := head
		for prev != tail {
			nex := p.Next
			p.Next = prev
			prev = p
			p = nex
		}
		return tail, head
	}

	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		tail := pre
		// 寻找尾部元素
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 如果不够说明不需要反转，直接返回
			if tail == nil {
				return dummy.Next
			}
		}

		// 反转前保存next
		nex := tail.Next
		// 反转这一段链表
		head, tail = myReverse(head, tail)

		// 虚拟结点的next指向新的head
		pre.Next = head
		// 新的tail指向旧的tail的next
		tail.Next = nex
		// pre指向上一段的tail
		pre = tail

		// 继续下一轮循环
		head = nex
	}

	return dummy.Next
}
