package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

删除链表的倒数第 N 个结点

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

*/
func main() {
	list1 := MakeListNode([]int{1, 2, 3, 5, 7})
	removeList1 := removeNthFromEnd1(list1, 3)
	PrintListNode(removeList1)

	list2 := MakeListNode([]int{1, 2, 3, 5, 7})
	removeList2 := removeNthFromEnd2(list2, 3)
	PrintListNode(removeList2)

	list3 := MakeListNode([]int{1, 2, 3, 5, 7})
	removeList3 := removeNthFromEnd3(list3, 3)
	PrintListNode(removeList3)
}

// removeNthFromEnd1 遍历两次,时间复杂度O(2*n),空间复杂度O(1)
// 第一次遍历获取链表长度，第二次遍历就知道是哪个位置要删除了，因为是倒数顺序
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var getLength func(head *ListNode) (length int)
	getLength = func(head *ListNode) (length int) {
		for ; head != nil; head = head.Next {
			length++
		}
		return
	}

	length := getLength(head)
	dummy := &ListNode{Next: head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// removeNthFromEnd2 辅助栈,时间复杂度O(n),空间复杂度O(n)
// 使用数组模拟也可以
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dummy := &ListNode{Next: head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// removeNthFromEnd3 双指针解法，最优解，时间复杂度O(n),空间复杂度O(1)
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	tail, slow := head, dummy
	// tail定位到要删除节点的位置，此时tail与slow距离是n+1
	for i := 0; i < n; i++ {
		if tail == nil {
			return head
		}

		tail = tail.Next
	}

	// tail继续移动直到末尾，同时slow也移动保持相对距离不变
	for ; tail != nil; tail = tail.Next {
		slow = slow.Next
	}

	// 此时slow就是要删除节点的前一个结点，也就是要改变Next指向的那个结点
	slow.Next = slow.Next.Next
	return dummy.Next
}
