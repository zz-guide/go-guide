package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/middle-of-the-linked-list/
链表的中间结点
*/
func main() {
	list1 := MakeListNode([]int{1, 2, 3, 4})
	removeList1 := middleNode1(list1)
	println("removeList1:", removeList1.Val)

	list2 := MakeListNode([]int{1, 2, 3, 4})
	removeList2 := middleNode2(list2)
	println("removeList2:", removeList2.Val)

	list3 := MakeListNode([]int{1, 2, 3, 4})
	removeList3 := middleNode3(list3)
	println("removeList3:", removeList3.Val)
}

// 快慢指针法，时间复杂度O(N),空间复杂度O(1)
func middleNode3(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

// 单指针法，两次遍历,时间复杂度O(2*N),空间复杂度O(1)
func middleNode2(head *ListNode) *ListNode {
	p, count := head, 0
	for p != nil {
		count++
		p = p.Next
	}

	p, count = head, count/2
	for count > 0 {
		p, count = p.Next, count-1
	}
	return p
}

// 使用一个数组存储,时间复杂度O(N),空间复杂度O(N)
func middleNode1(head *ListNode) *ListNode {
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	return arr[len(arr)/2]
}
