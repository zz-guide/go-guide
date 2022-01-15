package main

import (
	. "go-guide/datastruct/list/node"
)

/**
https://leetcode-cn.com/problems/merge-k-sorted-lists
合并K个有序链表
*/
func main() {
	// 两路归并，分治方式
	list1 := MakeListNode([]int{1, 2, 3, 5, 7})
	list2 := MakeListNode([]int{2, 4, 6})
	list3 := MakeListNode([]int{7, 9, 11})
	lists := []*ListNode{list1, list2, list3}
	mergeList := mergeKLists(lists)
	PrintListNode(mergeList)

}

func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	n := len(lists)
	for i := 0; i < n; i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		cur = lists[i]
		pre = mergeInBetween(pre, cur)
	}
	return pre
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
