package main

import (
	"fmt"
	. "go-guide/datastruct/list/node"
)

/**
https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
链表中倒数第k个节点
*/
func main() {
	list1 := MakeListNode([]int{1, 2, 3, 5, 7})
	searchList := getKthFromEnd(list1, 9)
	fmt.Println("searchList:", searchList.Val)
}

// 顺序查找，需要遍历2次
func getKthFromEnd1(head *ListNode, k int) (kth *ListNode) {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	for kth = head; n > k; n-- {
		kth = kth.Next
	}
	return
}

// 双指针，只遍历一次
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
