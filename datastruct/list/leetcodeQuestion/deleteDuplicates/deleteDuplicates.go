package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

删除排序链表中的重复元素

给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列

注意：1.list已排序，重复的元素必然挨着


*/
func main() {
	list1 := MakeListNode([]int{1, 2, 2, 3, 4})
	PrintListNode(deleteDuplicates(list1))
}

// deleteDuplicates 时间复杂度O(n),空间复杂度O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
