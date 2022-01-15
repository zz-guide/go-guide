package main

import (
	"container/heap"
	. "go-guide/datastruct/list/node"
)

/**
https://leetcode-cn.com/problems/merge-k-sorted-lists/
合并K个有序链表
*/
func main() {
	// 最小堆方式
	list1 := MakeListNode([]int{1, 2, 3, 5, 7})
	list2 := MakeListNode([]int{2, 4, 6})
	list3 := MakeListNode([]int{7, 9, 11})
	lists := []*ListNode{list1, list2, list3}
	mergeList := mergeKLists(lists)
	PrintListNode(mergeList)
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
}
