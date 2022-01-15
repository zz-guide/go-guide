package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/linked-list-cycle-ii/

环形链表 II

给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。

提示：

链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引


注意：1.不允许修改链表
2.如果允许修改的话还有其他方法

*/
func main() {
	list1 := MakeCircleListNode([]int{3, 2, 0, 4}, 1, 3)
	PrintSingleListNode(detectCycle(list1))
	PrintSingleListNode(detectCycle1(list1))
	//PrintSingleListNode(detectCycle2(list1))
	PrintSingleListNode(detectCycle3(list1))
}

// detectCycle 哈希
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// detectCycle1 快慢指针 时间复杂度O(N),空间复杂度O(1)
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 2(x+y) = (2-1)(z+y)+x+y
		// N>=2，2-1代表快指针需要多饶几圈，x=z,也就是slow和head同时移动x距离就是第一个相交结点的位置
		if slow == fast {
			tmp := head
			for slow != tmp {
				slow = slow.Next
				tmp = tmp.Next
			}

			return tmp
		}
	}

	return nil
}

// detectCycle2 3步也是可以的，但是会超时
func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		// 快指针走3步数，慢指针走1步数
		fast = fast.Next.Next.Next
		if slow == fast {
			tmp := head
			for slow != tmp {
				slow = slow.Next
				tmp = tmp.Next
			}

			return tmp
		}
	}

	return nil
}

// detectCycle3 把链表的每一个元素的Next都断掉了，不符合题意，最终会走一圈+x长度，碰到第一个Next == temp就是入口
func detectCycle3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	temp := &ListNode{Val: 0}
	cur := head
	for cur != nil {
		if cur.Next == temp {
			return cur
		}

		p := cur.Next
		cur.Next = temp
		cur = p
	}

	return nil
}
