package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/reverse-linked-list-ii/
反转链表 II

给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

提示：

链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n

进阶： 你可以使用一趟扫描完成反转吗？

*/
func main() {
	list1 := MakeListNode([]int{1, 2, 3, 5, 7})
	left := 2
	right := 4
	PrintListNode(reverseBetween(list1, left, right))
	list2 := MakeListNode([]int{1, 2, 3, 5, 7})
	PrintListNode(reverseBetween2(list2, left, right))
}

// reverseBetween O(n) O(1) 直接反转
func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Next: head}

	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// pre = left之前的一个结点，反转的过程中pre的指向需要一直变化
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		// 下一个新的头结点
		next := cur.Next
		// 当前结点指向下一个的下一个
		cur.Next = next.Next
		// 新的头结点指向old的头结点
		next.Next = pre.Next
		// pre指向新的头结点
		pre.Next = next
	}

	return dummyNode.Next
}

func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		// 临时保存下一个值，准备移动
		next := cur.Next
		// 更改指向上一个值
		cur.Next = pre
		// 标记当前值为上一个值
		pre = cur
		// 移动到下一个位置
		cur = next
	}
}

// reverseBetween2 O(n) O(1) 先切断，再反转，然后链接
func reverseBetween2(head *ListNode, left, right int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	reverseLinkedList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}
