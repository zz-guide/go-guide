package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目: https://leetcode-cn.com/problems/swap-nodes-in-pairs/
两两交换链表中的节点

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

输入：head = [1,2,3,4]
输出：[2,1,4,3]

提示：

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100

*/
func main() {
	head := MakeListNode([]int{1, 2, 3, 5, 7})
	head1 := MakeListNode([]int{1, 2, 3, 5, 7})
	PrintListNode(swapPairs1(head))
	PrintListNode(swapPairs2(head1))
}

// swapPairs2 递归
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs2(newHead.Next)
	newHead.Next = head
	return newHead
}

// swapPairs1 迭代
func swapPairs1(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	temp := dummyHead
	// 虚拟头结点->1->2->3->4
	// cur->node1->node2->node3
	// 第一步: cur->node2
	// 第二步: node1->node3
	// 第三步: node2->node1， cur->node2->node1->node3
	// 第四步: 虚拟头结点移动到node1位置
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
