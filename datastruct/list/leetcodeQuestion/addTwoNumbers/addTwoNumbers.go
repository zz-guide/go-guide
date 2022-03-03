package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/add-two-numbers/

两数相加

给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以0开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零


注意：
	1.每个结点大小都是0~9，存在进位
	2.不管是什么顺序，数字相加都是从低位向高位加，不够补0，跟正序逆序没关系

*/
func main() {
	list1 := MakeListNode([]int{2, 3, 4})
	list2 := MakeListNode([]int{5, 6, 4})
	PrintListNode(addTwoNumbers(list1, list2))
}

// addTwoNumbers 直接模拟, O(max(m,n)),O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	cur := head
	var carry int // 进位
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return head.Next
}
