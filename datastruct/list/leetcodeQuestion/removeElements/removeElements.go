package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/remove-linked-list-elements/
移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

提示：

列表中的节点数目在范围 [0, 104] 内
1 <= Node.val <= 50
0 <= val <= 50

*/
func main() {
	list1 := MakeListNode([]int{1, 2, 6, 3, 4, 5, 6})
	val := 6
	removeList1 := removeElements(list1, val)
	PrintListNode(removeList1)

	list2 := MakeListNode([]int{1, 2, 6, 3, 4, 5, 6})
	removeList2 := removeElements1(list2, val)
	PrintListNode(removeList2)
}

// removeElements 迭代
func removeElements(head *ListNode, val int) *ListNode {
	var virtualNode = &ListNode{Next: head}

	cur := virtualNode
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return virtualNode.Next
}

// removeElements1 递归
func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
