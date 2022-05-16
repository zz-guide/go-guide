package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/partition-list/

分隔链表

给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。


提示：

链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200

注意:
	1.不能交换位置，因为要相对顺序不能变化

*/

func main() {
	list1 := MakeListNode([]int{1, 4, 3, 2, 5, 2})
	x := 3
	PrintListNode(partition(list1, x))
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small

	large := &ListNode{}
	largeHead := large

	// 只需要顺序遍历，把小的值按顺序相连接，大的连接即可
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}

		head = head.Next
	}

	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
