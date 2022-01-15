package main

import (
	. "go-guide/datastruct/list/node"
)

/**
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists/
合并2个有序链表
*/
func main() {
	// 迭代方式
	list1 := MakeListNode([]int{1, 2, 3, 5, 7})
	list2 := MakeListNode([]int{2, 4, 6})
	mergeList := mergeInBetween(list1, list2)
	PrintListNode(mergeList)

	// 递归方式
	list3 := MakeListNode([]int{1, 2, 3, 5, 7})
	list4 := MakeListNode([]int{2, 4, 6})
	mergeList1 := mergeInBetween1(list3, list4)
	PrintListNode(mergeList1)
}

func mergeInBetween(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	//设立一个虚拟结点
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	// 此时list还没遍历完毕，list2已全部遍历完毕，list1剩余结点不需要继续遍历，因为是有序的，可以直接加到curr的Next
	if list1 != nil {
		curr.Next = list1
	}

	// list2同样
	if list2 != nil {
		curr.Next = list2
	}

	return dummy.Next
}

// mergeTwoLists1 递归方式
func mergeInBetween1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeInBetween1(list1.Next, list2)
		return list1
	}

	list2.Next = mergeInBetween1(list1, list2.Next)
	return list2
}
