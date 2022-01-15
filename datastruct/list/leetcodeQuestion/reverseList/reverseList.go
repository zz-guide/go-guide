package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/reverse-linked-list/
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

提示：链表中节点的数目范围是 [0, 5000]	-5000 <= Node.val <= 5000
进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000

解法：for循环遍历，递归，栈三种（反转数组，反转字符串这种的都可以使用栈来解决，就是比较浪费空间）
*/
func main() {
	head := MakeListNode([]int{1, 2, 3, 5, 7})
	PrintListNode(reverseList(head))
}

// reverseList 时间复杂度：O(n)，其中 nn 是链表的长度。需要遍历链表一次。空间复杂度：O(1)
func reverseList(head *ListNode) *ListNode {
	var newHead, next *ListNode
	curr := head
	for curr != nil {
		next = curr.Next
		// 反转
		curr.Next = newHead
		// 替换成为新的头结点
		newHead = curr
		// 继续向下迭代
		curr = next
	}
	return newHead
}

// reverseList1 递归 时间复杂度：O(n)，空间复杂度：O(n)
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList1(head.Next)
	// 翻转头节点与第二个节点的指向
	head.Next.Next = head
	// 此时的 head 节点为尾节点，next 需要指向 NULL
	head.Next = nil
	return newHead
}
