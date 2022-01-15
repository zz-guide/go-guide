package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/submissions/
判断两个链表是否相交

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

提示：

listA 中节点数目为 m
listB 中节点数目为 n
0 <= m, n <= 3 * 104
1 <= Node.val <= 105
0 <= skipA <= m
0 <= skipB <= n
如果 listA 和 listB 没有交点，intersectVal 为 0
如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1]


注意：1.题目要求找到两个链表第一个相交的结点，返回即可
2.没有环，都是单链表
3.相交是结点必须是地址一样

*/
func main() {
	intersectionNode := MakeListNode([]int{8, 4, 5})
	list1 := MakeListNodeByNode([]int{4, 1}, intersectionNode)
	list2 := MakeListNodeByNode([]int{5, 0, 1}, intersectionNode)
	//PrintListNode(getIntersectionNode(list1, list2))
	//PrintListNode(getIntersectionNode1(list1, list2))
	PrintListNode(getIntersectionNode2(list1, list2))
}

// 哈希 时间复杂度：O(m+n) 空间复杂度:O(m)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		// 返回第一个相交的结点即可
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

// 双指针 时间复杂度：O(m+n) 空间复杂度:O(1)
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}

		//fmt.Println("pa:", pa.Val)
		//fmt.Println("pb:", pb.Val)
	}

	return pa
}

// getIntersectionNode2 利用长度
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
