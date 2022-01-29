package main

import (
	"container/list"
	. "go-guide/datastruct/list/node"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

限制：
0 <= 链表长度 <= 10000

*/
func main() {
	head := MakeListNode([]int{1, 2, 3, 5, 7})
	head1 := MakeListNode([]int{1, 2, 3, 5, 7})
	head2 := MakeListNode([]int{1, 2, 3, 5, 7})
	head3 := MakeListNode([]int{1, 2, 3, 5, 7})
	log.Println("从尾到头打印链表-递归法", reversePrint(head))
	log.Println("从尾到头打印链表-反转链表", reversePrint1(head1))
	log.Println("从尾到头打印链表-反转数组", reversePrint2(head2))
	log.Println("从尾到头打印链表-栈", reversePrint3(head3))
}

// reversePrint 递归法，从尾部开始append即可,效率不高
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var appendData func(node *ListNode) []int
	appendData = func(node *ListNode) []int {
		if node.Next != nil {
			list := appendData(node.Next)
			list = append(list, node.Val)
			return list
		}

		return []int{node.Val}
	}

	return appendData(head)
}

// reversePrint1 反转链表然后打印
func reversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}

	cur := head
	// 反转
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next, pre = pre, cur
		cur = next
	}

	// 遍历
	var res []int
	for pre != nil {
		res = append(res, pre.Val)
		pre = pre.Next
	}

	return res
}

// reversePrint2 先遍历链表，然后反转数组
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var res []int
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

// reversePrint3 利用栈
func reversePrint3(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := list.New()
	for head != nil {
		res.PushFront(head.Val)
		head = head.Next
	}

	var ret []int
	for e := res.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(int))
	}

	return ret
}
