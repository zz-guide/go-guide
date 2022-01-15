package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/linked-list-cycle/

环形链表

给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。


*/
func main() {
	list1 := MakeListNode([]int{1, 2, 3, 4})
	bool1 := hasCycle(list1)
	println("环形链表-哈希:", bool1)

	list2 := MakeListNode([]int{1, 2, 3, 4})
	bool2 := hasCycle1(list2)
	println("环形链表-双指针:", bool2)
}

// 哈希表 时间复杂度O(N) 空间复杂度O(N)
func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 快慢指针 时间复杂度O(N),空间复杂度O(1)
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
